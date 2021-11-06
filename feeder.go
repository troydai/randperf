package randperf

import (
	"math/rand"
	"sync"
)

// feeder pattern

var (
	feeder     <-chan float64
	feederOnce *sync.Once = &sync.Once{}
)

func NewFeederGen() Generator {
	feederOnce.Do(newRandFloat64Feeder)

	return func() float64 {
		return <-feeder
	}
}

func newRandFloat64Feeder() {
	output := make(chan float64)
	go func() {
		for {
			output <- rand.Float64()
		}
	}()

	feeder = output
}
