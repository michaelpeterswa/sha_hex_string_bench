# sha_hex_string_bench

Comparing to see if `fmt.Sprintf()` is faster than `hex.EncodeToString()`. Newsflash, it's not. Reflects are slow.

## Results
```
λ  sha_hex_string_bench main ✗  go test -run=XXX -bench=. -benchtime=100000x
goos: darwin
goarch: amd64
pkg: github.com/michaelpeterswa/sha_hex_string_bench
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkShaHexStringWithHexLib-12     	  100000	       729.5 ns/op	     176 B/op	       4 allocs/op
BenchmarkShaHexStringWithSprintf-12    	  100000	       966.3 ns/op	     176 B/op	       5 allocs/op
BenchmarkRandStringBytes-12            	  100000	       382.2 ns/op	      48 B/op	       2 allocs/op
PASS
ok  	github.com/michaelpeterswa/sha_hex_string_bench	0.336s
```
