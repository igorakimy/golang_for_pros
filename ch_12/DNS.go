package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	// Получить список имен, которые соответствуют переданному IP-адресу.
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookupHostname(hostname string) ([]string, error) {
	// Получить список IP-адресов, которые соответствуют переданному хосту.
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	input := arguments[1]
	// Выполнить синтаксический анализ строки на наличие IPv4 или IPv6 адреса.
	IPAddress := net.ParseIP(input)

	if IPAddress == nil {
		IPs, err := lookupHostname(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println(singleIP)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}
