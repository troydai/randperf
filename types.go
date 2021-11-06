package randperf

import (
	"math/rand"
	"time"
)

type Generator func() float64

func noopWork() Generator {
	return func() float64 {
		return 0
	}
}

func smallWork() Generator {
	return func() float64 {
		time.Sleep(time.Millisecond)
		return 0
	}
}

func getRand() Generator {
	return func() float64 {
		return rand.Float64()
	}
}
