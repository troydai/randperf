package randperf

type signalChan <-chan struct{}

func workChan(
	done <-chan struct{},
	output chan<- float64,
	generator func() float64) {
	go func() {
		for {
			select {
			case <-done:
				return
			case output <- generator():
			}
		}
	}()
}
