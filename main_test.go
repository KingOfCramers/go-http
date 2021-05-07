package main

import (
	"testing"
)

func BenchmarkGiveMeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		giveMeSeven()
	}
}

func TestGiveMeSeven(t *testing.T) {
	res := giveMeSeven()
	if res != 7 {
		t.Errorf("Expected 7 but got %d", res)
	}
}
