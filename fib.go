package fib

import (
	"math/big"
)

const (
	zero = 0
)

var (
	bigOne = big.NewInt(1)
	bigTwo = big.NewInt(2)
)

// Fib returns fibonacci number (int) for index n (int) with a simple loop algorithm
func Fib(n int) int {
	f0 := 0
	f1 := 1
	for i := 0; i < n - 1; i++ {
		f1 = f0 + f1
		f0 = f1 - f0
	}
	return f1
}

// BigFib returns fibonacci number (big.Int) for index n (big.Int) with a single loop algorithm
// this method supports very large numbers as input and output
func BigFib(n *big.Int) *big.Int {
	n = n.Sub(n, bigOne)
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(n) < 0; i.Add(i, bigOne) {
		s := f0.Add(f1, f0)
		f0 = f1
		f1 = s
	}
	return f1
}

// RFib returns fibonacci number (int) for index n (int) with recursive function algorithm
func RFib(n int) int {
	k, l := rfib(n - 1)
	return k+l
}

func rfib(n int) (int, int) {
	if  n < 2 {
		return zero, n
	}
	k, l := rfib(n - 1)
	return l, k + l
}

// RBigFib returns fibonacci number (big.Int) for index n (big.Int) with recursive function algorithm
// this method supports very large numbers as input and output
func RBigFib(n *big.Int) *big.Int {
	k, l := rbfib(n.Sub(n, bigOne))
	return k.Add(k, l)
}

func rbfib(n *big.Int) (*big.Int, *big.Int) {
	if  n.Cmp(bigTwo) < 0 {
		return big.NewInt(0), n
	}
	k, l := rbfib(n.Sub(n, bigOne))
	return l, k.Add(k, l)
}