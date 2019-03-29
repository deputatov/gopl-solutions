package main

import (
	"testing"
)

const (
	iterationsNum = 100000
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < iterationsNum; i++ {
		Echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < iterationsNum; i++ {
		Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < iterationsNum; i++ {
		Echo3()
	}
}

/*
0.113s
0.136s
0.050s
*/