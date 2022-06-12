package fibonacci

import (
	"math/big"
)

func Fib(n uint) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for i := uint(1); i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}
