package main

import "testing"

var num = 1000

func BenchmarkPrimeNumber(b *testing.B) { // B represent the Benchmark struct
	for i := 0; i < b.N; i++ { // N = No. of iterations will dynamically decided by our benchmark
		PrimeNumber(num)
	}
}
