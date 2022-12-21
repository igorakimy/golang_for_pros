package main

import (
	"fmt"
	"time"
)

const (
	SUN time.Weekday = iota
	MON
	TUE
	WEN
	THU
	FRI
	SAT
)

func main() {
	fmt.Println(MON, TUE, WEN, THU, FRI, SAT, SUN)
}
