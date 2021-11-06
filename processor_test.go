package randperf

import (
	"fmt"
	"sync"
	"testing"
)

var workerNums = []int{1, 10, 100, 1000}
var workLoad = 1000

func BenchmarkProcessNoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		procSync(workLoad, noopWork)
	}
}

func BenchmarkProcessAsyncNoop(b *testing.B) {
	for _, workerNum := range workerNums {
		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			benchmarkRoutines(b, workerNum, workLoad, noopWork)
		})
	}
}

func BenchmarkProcessSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		procSync(workLoad, smallWork)
	}
}

func BenchmarkProcessAsyncSmall(b *testing.B) {
	for _, workerNum := range workerNums {
		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			benchmarkRoutines(b, workerNum, workLoad, smallWork)
		})
	}
}

func BenchmarkProcessRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		procSync(workLoad, getRand)
	}
}

func BenchmarkProcessAsyncRand(b *testing.B) {
	for _, workerNum := range workerNums {
		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			benchmarkRoutines(b, workerNum, workLoad, getRand)
		})
	}
}

func BenchmarkProcessAsyncRandInd(b *testing.B) {
	for _, workerNum := range workerNums {
		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			benchmarkRoutines(b, workerNum, workLoad, NewPrebuildRand)
		})
	}
}

func BenchmarkProcessAsyncRandFeeder(b *testing.B) {
	for _, workerNum := range workerNums {
		b.Run(fmt.Sprintf("worker_%v", workerNum), func(b *testing.B) {
			benchmarkRoutines(b, workerNum, workLoad, NewFeederGen)
		})
	}
}

func benchmarkRoutines(b *testing.B, workerNum, workAmount int, gen func() Generator) {
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(workerNum)
		for j := 0; j < workerNum; j++ {
			procRoutine(workAmount, gen(), wg)
		}
		wg.Wait()
	}
}
