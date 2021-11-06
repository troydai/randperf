package randperf

import (
	"math/rand"
	"sync"
)

const prebuildSize = 5000

var initOnce = &sync.Once{}
var pre []float64

func prebuild() {
	initOnce.Do(func() {
		for i := 0; i < prebuildSize; i++ {
			pre = append(pre, rand.Float64())
		}
	})
}

func NewPrebuildRand() Generator {
	prebuild()

	at := rand.Intn(len(pre))
	return func() float64 {
		v := pre[at%len(pre)]
		at++
		return v
	}
}
