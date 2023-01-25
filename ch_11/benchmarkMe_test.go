package main

import (
	"testing"
)

var result int

func benchmark_fibonacci1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci1(n)
	}
	result = r
}

func benchmark_fibonacci2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci2(n)
	}
	result = r
}

func benchmark_fibonacci3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci3(n)
	}
	result = r
}

func Benchmark30Fibonacci1(b *testing.B) {
	benchmark_fibonacci1(b, 30)
}

func Benchmark30Fibonacci2(b *testing.B) {
	benchmark_fibonacci2(b, 30)
}

func Benchmark30Fibonacci3(b *testing.B) {
	benchmark_fibonacci3(b, 30)
}

func Benchmark50Fibonacci1(b *testing.B) {
	benchmark_fibonacci1(b, 50)
}

func Benchmark50Fibonacci2(b *testing.B) {
	benchmark_fibonacci2(b, 50)
}

func Benchmark50Fibonacci3(b *testing.B) {
	benchmark_fibonacci3(b, 50)
}
