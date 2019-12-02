package main

import (
	"fmt"
	"sync"
	"time"
)

func computeSeries(start, end int64) (sum float64) {
	increment := float64(2*end - 1)
	sign := 1.0
	if end%2 == 0 {
		sign = -1.0
	}
	for i := end; i >= start; i-- {
		sum += 1 / (increment * sign)
		sign *= -1
		increment -= 2.0
	}
	return
}

func main() {
	var N int64 = 1_000_000_000
	// Computing the series sequentially
	start := time.Now()
	fmt.Println("The approx value of pi: ", 4*computeSeries(1, N))
	fmt.Println("Time elapsed: ", time.Now().Sub(start))

	// With go routines
	start2 := time.Now()
	var wg sync.WaitGroup
	noRoutines := 4
	wg.Add(noRoutines)
	var values [4]float64

	go func() {
		values[0] = computeSeries(1, 250_000_000)
		wg.Done()
	}()
	go func() {
		values[1] = computeSeries(250_000_001, 500_000_000)
		wg.Done()
	}()
	go func() {
		values[2] = computeSeries(500_000_001, 750_000_000)
		wg.Done()
	}()
	go func() {
		values[3] = computeSeries(750_000_001, 1_000_000_000)
		wg.Done()
	}()
	wg.Wait()

	sum := 0.0
	for i := 0; i < noRoutines; i++ {
		sum += values[i]
	}
	fmt.Println("The approx value of pi: ", sum*4)
	fmt.Println("Time elapsed: ", time.Now().Sub(start2))
}
