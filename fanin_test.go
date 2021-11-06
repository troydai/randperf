package randperf

import (
	"fmt"
	"testing"
)

func BenchmarkFanInNoopWork(b *testing.B) {
	gen := noopWork()
	for i := 0; i < b.N; i++ {
		gen()
	}
}

func BenchmarkFanInNoopWorker(b *testing.B) { benchmarkWorker(b, noopWork) }

func BenchmarkFanInSmallWork(b *testing.B) {
	gen := smallWork()
	for i := 0; i < b.N; i++ {
		gen()
	}
}

func BenchmarkFanInSmallWorker(b *testing.B) { benchmarkWorker(b, smallWork) }

func BenchmarkFanInRand(b *testing.B) {
	gen := getRand()
	for i := 0; i < b.N; i++ {
		gen()
	}
}

func BenchmarkFanInRandWorker(b *testing.B)         { benchmarkWorker(b, getRand) }
func BenchmarkFanInPrebuiltRandWorker(b *testing.B) { benchmarkWorker(b, NewPrebuildRand) }

func benchmarkWorker(b *testing.B, getGen func() Generator) {
	output := make(chan float64)
	done := make(chan struct{})
	defer close(done)

	for _, workerNum := range workerNums {
		for i := 0; i < workerNum; i++ {
			workChan(done, output, getGen())
		}

		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				<-output
			}
		})
	}
}
