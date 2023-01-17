package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var maxSysCalls = 0

const SYSCALLFILE = "SYSCALLS"

func main() {
	var SYSTEMCALLS []string
	f, err := os.Open(SYSCALLFILE)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "SYS_", "", -1)
		temp := strings.ToLower(strings.Split(line, "-")[0])
		SYSTEMCALLS = append(SYSTEMCALLS, temp)
		maxSysCalls++
	}

	// В срезе COUNTER хранятся данные о том, сколько раз каждый системный вызов
	// был обнаружен в отслеживаемой программе.
	COUNTER := make([]int, maxSysCalls)
	var regs syscall.PtraceRegs
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Wait:", err)
	}

	pid := cmd.Process.Pid
	fmt.Println("Process ID:", pid)

	before := true
	forCount := 0
	for {
		if before {
			err := syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			if regs.Orig_rax > uint64(maxSysCalls) {
				fmt.Println("Unknown:", regs.Orig_rax)
				return
			}

			COUNTER[regs.Orig_rax]++
			forCount++
		}

		// Эта функция заставляет Go продолжать выполнение отслеживаемой программы
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			fmt.Println("PtraceSyscall:", err)
			return
		}

		// Но останавливаться, когда программа входит или выходит из системного вызова.
		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			fmt.Println("Wait4:", err)
			return
		}
		before = !before
	}

	for i, x := range COUNTER {
		if x != 0 {
			fmt.Println(SYSTEMCALLS[i], "->", x)
		}
	}
	fmt.Println("Total System Calls:", forCount)
}
