package main

import "testing"

func BenchmarkTestAllocationForLoop(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TestAllocationForLoop()
	}
}

func BenchmarkTestExponentialCopyPattern(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TestAllocationExponentialCopyPattern()
	}
}
