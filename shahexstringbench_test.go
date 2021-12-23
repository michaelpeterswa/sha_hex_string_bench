package shahexstringbench

// go test -run=XXX -bench=. -benchtime=100000x

import "testing"

const strlen = 20

// from shahextstringbench_test.go
func BenchmarkShaHexStringWithHexLib(b *testing.B) {
	b.ReportAllocs()	
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ShaHexStringWithHexLib(RandStringBytes(strlen))
	}
}

// from shahextstringbench_test.go
func BenchmarkShaHexStringWithSprintf(b *testing.B) {
	b.ReportAllocs()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ShaHexStringWithSprintf(RandStringBytes(strlen))
	}
}

func BenchmarkRandStringBytes(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		RandStringBytes(strlen)
	}
}
