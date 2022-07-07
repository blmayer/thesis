package generator

import (
	"math/big"
)

// NumOfDivisors calculates the number of divisors of a given number
func NumOfDivisors(n *big.Int) int {
	num := 2
	temp := big.NewInt(0)
	for i := big.NewInt(2); i.Cmp(n) == -1; i.Add(i, big.NewInt(1)) {
		if temp.Rem(n, i).Cmp(big.NewInt(0)) == 0 {
			num++
		}
	}

	return num
}
