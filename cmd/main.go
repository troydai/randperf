package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/troydai/randperf"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	test("math/rand", 1000000, rand.Float64)

	gen := randperf.NewPrebuildRand()
	test("Prebuild Rand", 1000000, gen)

	done := make(chan struct{})
	defer close(done)
	test("Prebuild Rand / Multiworker", 1000000, merge(16, done))
}

func test(name string, sampleSize int, gen randperf.Generator) {
	var samples []float64
	for i := 0; i < 1000000; i++ {
		samples = append(samples, gen())
	}

	mean, err := stats.Mean(samples)
	if err != nil {
		log.Fatal(err)
	}

	stdp, err := stats.StdDevP(samples)
	if err != nil {
		log.Fatal(err)
	}

	stds, err := stats.StdDevS(samples)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(name)
	fmt.Printf("Mean: %v\n", mean)
	fmt.Printf("Stdp: %v\n", stdp)
	fmt.Printf("Stds: %v\n", stds)
}

func merge(workerNum int, done <-chan struct{}) randperf.Generator {
	output := make(chan float64)
	for i := 0; i < workerNum; i++ {
		go func() {
			localGen := randperf.NewPrebuildRand()
			for {
				select {
				case <-done:
					return
				case output <- localGen():
				}
			}
		}()
	}

	return func() float64 {
		return <-output
	}
}
