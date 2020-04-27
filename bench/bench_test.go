package main

import (
	"fabnic"
	"testing"
)

func BenchmarkBenchFab2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fabnic.Fab2(40)
	}
}
func BenchmarkBenchFab1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fabnic.Fab1(40)
	}
}
