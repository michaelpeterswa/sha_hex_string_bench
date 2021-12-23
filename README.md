# sha_hex_string_bench

Comparing to see if `fmt.Sprintf()` is faster than `hex.EncodeToString()`. Newsflash, it's not. Reflects are slow.

## Results
```
λ  ~/go/src/github.com/michaelpeterswa/sha_hex_string_bench  go test -run=XXX -bench=. -benchtime=100000x
goos: darwin
goarch: amd64
pkg: github.com/michaelpeterswa/sha_hex_string_bench
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkShaHexStringWithHexLib-12     	  100000	       706.9 ns/op
BenchmarkShaHexStringWithSprintf-12    	  100000	      1007 ns/op
BenchmarkRandStringBytes-12            	  100000	       400.3 ns/op
PASS
ok  	github.com/michaelpeterswa/sha_hex_string_bench	0.305s
```
