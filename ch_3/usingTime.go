package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	// Epoch time: 1671344935
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	// 2022-12-18 09:28:55.9293254 +0300 MSK m=+0.002091801 2022-12-18T09:28:55+03:00
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())
	// Sunday 18 December 2022

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))
	// Time difference: 1.0342295s

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	// 12 December 2022
	loc, _ := time.LoadLocation("Europe/Moscow")
	moscowTime := t.In(loc)
	fmt.Println("Moscow:", moscowTime)
	// Moscow: 2022-12-18 09:28:55.9293254 +0300 MSK
}
