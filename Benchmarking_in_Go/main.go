package main

import (
	"fmt"
	"math"
)

// What is benchmarking in golang ?
// Each output against a standard , assessing the code's overall performance level

// Testing is built-in provided by Golang and we can do the benchmarking by testing
// _test.go // suffix

// Way of Bench Marking
// func Benchmark<Func-name>(b *testing.B){}  -->>> b pointer reperesent the testing.B<Benchmark>

//
// ------------------- BenchMark Command ---------------------
// go test -bench=.   // --->>> =. Means run all your benchmark files in your current path
// go test -bench=. -count 10  // ---->> Benchmark the code till 10 times
//

func main() {
	fmt.Println(PrimeNumber(10))
}

func PrimeNumber(max int) []int {
	var primes []int

	for i := 0; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ { //   Only go up to the square root of n
			fmt.Println("The Math Value is", int(math.Sqrt(float64(i))))
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
