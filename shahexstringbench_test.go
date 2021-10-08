package shahexstringbench

import "testing"

const strlen = 20

// from shahextstringbench_test.go
func BenchmarkShaHexStringWithHexLib(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ShaHexStringWithHexLib(RandStringBytes(strlen))
	}
}

// from shahextstringbench_test.go
func BenchmarkShaHexStringWithSprintf(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ShaHexStringWithSprintf(RandStringBytes(strlen))
	}
}

func BenchmarkRandStringBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandStringBytes(strlen)
	}
}
