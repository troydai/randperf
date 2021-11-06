package randperf

import "sync"

func procSync(count int, gen func() Generator) {
	work := gen()
	for i := 0; i < count; i++ {
		work()
	}
}

func procRoutine(count int, gen Generator, finished *sync.WaitGroup) {
	go func() {
		for i := 0; i < count; i++ {
			gen()
		}
		finished.Done()
	}()
}
