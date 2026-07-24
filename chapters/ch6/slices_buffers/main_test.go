package main

import "testing"

func testOpenResource(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := OpenResource()
		if err != nil {
			return
		}
	}
}
