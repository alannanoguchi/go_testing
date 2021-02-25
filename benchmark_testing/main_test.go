package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

// Checking to see if there are any performance differences between different inputs
func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }  // Benchmark 1
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) } // Benchmark 2
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }   // Benchmark 3

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}
