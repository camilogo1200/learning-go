package main

import "testing"

/*
benchmark:
go test -bench=. -benchmem

1. Generate the profile files:
go test -bench=. -benchmem -memprofile=mem.out -cpuprofile=cpu.out

2. Analyze via the interactive browser UI
go tool pprof -http=:8080 mem.out

Collect and analyze the heap while your application is under load
go tool pprof <executable / endpoint> -> go tool pprof http://localhost:6060/debug/pprof/heap

repeat the test for accuracy -> go test -bench=. -count=5

more info: https://go.dev/doc/diagnostics
*/

func testOpenResource(b *testing.B) {

	//b.ResetTimer() if you have heavy setup code that you do not want included in the performance score.
	b.ReportAllocs() // Explicitly report memory for this benchmark
	for i := 0; i < b.N; i++ {
		_, err := OpenResource()
		if err != nil {
			return
		}
	}
}
