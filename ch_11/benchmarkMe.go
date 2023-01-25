package main

import (
	"fmt"
)

func fibonacci1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci1(n-1) + fibonacci1(n-2)
	}
}

func fibonacci2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci2(n-1) + fibonacci2(n-2)
}

func fibonacci3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fibonacci1(40))
	fmt.Println(fibonacci2(40))
	fmt.Println(fibonacci3(40))
}
