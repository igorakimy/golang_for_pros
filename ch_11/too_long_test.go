package main

import (
	"testing"
)

func Test_test_one(t *testing.T) {
	sleepWithMe()
	value := getOne()
	if value != 1 {
		t.Errorf("Function returned %v", value)
	}
	sleepWithMe()
}

func Test_test_two(t *testing.T) {
	sleepWithMe()
	value := getTwo()
	if value != 2 {
		t.Errorf("Funtion returned %v", value)
	}
}

func Test_that_will_fail(t *testing.T) {
	value := getOne()
	if value != 2 {
		t.Errorf("Function returned %v", value)
	}
}
