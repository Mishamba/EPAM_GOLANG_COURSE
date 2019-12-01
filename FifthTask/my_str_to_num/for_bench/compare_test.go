package main

import (
	"testing"
)

func BenchmarkMyStrToInt1(b *testing.B) {
	someString := []string{
		"1234",
		"4321",
		"3212",
		"643133",
	}
	for i := 0; i < b.N; i++ {
		_, _ = myStrToNum1(someString[i%len(someString)], "int")
	}
}

func BenchmarkByStrToInt2(b *testing.B) {
	someString := []string{
		"1234",
		"4321",
		"3212",
		"643133",
	}
	for i := 0; i < b.N; i++ {
		_, _ = myStrToNum2(someString[i%len(someString)], "int")
	}
}
