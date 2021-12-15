package fib

import (
	"math/big"
	"testing"
)

func BenchmarkFib90(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(90)
	}
}

func BenchmarkRFib90(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RFib(90)
	}
}

func BenchmarkBigFib300(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BigFib(big.NewInt(300))
	}
}

func BenchmarkRBigFib300(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RBigFib(big.NewInt(300))
	}
}