# Performance of rand in concurrency

```
➜  randperf git:(master) ✗ go test -bench=. -cpu=1 -benchtime=2s                     
goos: darwin
goarch: amd64
pkg: github.com/troydai/randperf
cpu: Intel(R) Core(TM) i3-1000NG4 CPU @ 1.10GHz
BenchmarkProcessNoop                     1860734              1284 ns/op
BenchmarkProcessAsyncNoop/worker_1               1381147              1734 ns/op
BenchmarkProcessAsyncNoop/worker_10               150769             15513 ns/op
BenchmarkProcessAsyncNoop/worker_100               15363            156536 ns/op
BenchmarkProcessAsyncNoop/worker_1000               1436           1610983 ns/op
BenchmarkProcessSmall                                  2        1208148756 ns/op
BenchmarkProcessAsyncSmall/worker_1                    2        1221858904 ns/op
BenchmarkProcessAsyncSmall/worker_10                   2        1221623902 ns/op
BenchmarkProcessAsyncSmall/worker_100                  2        1228655956 ns/op
BenchmarkProcessAsyncSmall/worker_1000                 2        1139762531 ns/op
BenchmarkProcessRand                              129150             18317 ns/op
BenchmarkProcessAsyncRand/worker_1                126506             18879 ns/op
BenchmarkProcessAsyncRand/worker_10                12860            186104 ns/op
BenchmarkProcessAsyncRand/worker_100                1268           1869062 ns/op
BenchmarkProcessAsyncRand/worker_1000                126          18767834 ns/op
BenchmarkProcessAsyncRandInd/worker_1             547615              3730 ns/op
BenchmarkProcessAsyncRandInd/worker_10             67917             35438 ns/op
BenchmarkProcessAsyncRandInd/worker_100             6428            352489 ns/op
BenchmarkProcessAsyncRandInd/worker_1000             664           3588374 ns/op
BenchmarkFanInNoopWork                          1000000000               0.3160 ns/op
BenchmarkFanInNoopWorker/worker_1                9860506               241.7 ns/op
BenchmarkFanInNoopWorker/worker_10               9681026               250.4 ns/op
BenchmarkFanInNoopWorker/worker_100              9740733               243.6 ns/op
BenchmarkFanInNoopWorker/worker_1000             9815878               244.2 ns/op
BenchmarkFanInSmallWork                             2048           1206564 ns/op
BenchmarkFanInSmallWorker/worker_1                  2019           1224514 ns/op
BenchmarkFanInSmallWorker/worker_10                21564            110243 ns/op
BenchmarkFanInSmallWorker/worker_100              236442             10350 ns/op
BenchmarkFanInSmallWorker/worker_1000            1942803              1171 ns/op
BenchmarkFanInRand                              137662302               17.48 ns/op
BenchmarkFanInRandWorker/worker_1                9248401               253.8 ns/op
BenchmarkFanInRandWorker/worker_10               9368907               255.5 ns/op
BenchmarkFanInRandWorker/worker_100              9372254               257.7 ns/op
BenchmarkFanInRandWorker/worker_1000             9336038               255.5 ns/op
BenchmarkFanInPrebuiltRandWorker/worker_1        9777374               243.3 ns/op
BenchmarkFanInPrebuiltRandWorker/worker_10       9225178               244.0 ns/op
BenchmarkFanInPrebuiltRandWorker/worker_100              9775420               245.5 ns/op
BenchmarkFanInPrebuiltRandWorker/worker_1000             9757562               245.5 ns/op
```