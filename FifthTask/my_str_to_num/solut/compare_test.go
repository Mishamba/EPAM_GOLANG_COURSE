package main

import (
	"strconv"
	"testing"
)

func BenchmarkMyStrToNum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myStrToNum1(strconv.Itoa(i))
	}
}

func BenchmarkMyStrToNum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myStrToNum2(strconv.Itoa(i))
	}
}
