package visibility

import (
	"math/big"
)

// Natural implements the natural visibility algorithm
func Natural(ts [10000]big.Int, i, j int) bool {
	// ts[x] >= ts[a]+(ts[b]-ts[a]) * (x-a)/(b-a)
	for x := i + 1; x < j; x++ {
		// ts[b] - ts[a]
		sub := big.NewInt(0)
		sub.Sub(&ts[j], &ts[i])
		subconv := big.NewFloat(0.0)
		subconv.SetInt(sub)

		// (x-a) / (b-a)
		frac := big.NewFloat(float64(x-i) / float64(j-i))

		// (ts[b] - ts[a]) * (x-a) / (b-a)
		mul := big.NewFloat(0.0)
		mul.Mul(subconv, frac)

		// ts[a] + (ts[b] - ts[a]) * (x-a) / (b-a)
		tsconv := big.NewFloat(0.0)
		tsconv.SetInt(&ts[i])
		res := big.NewFloat(0.0)
		res.Add(tsconv, mul)

		// Comparison
		// ts[x] >= res
		txconv := big.NewFloat(0.0)
		txconv.SetInt(&ts[x])
		if txconv.Cmp(res) != -1 {
			return false
		}
	}
	return true
}

// Horizontal implements the horizontal visibility algorithm
func Horizontal(ts [10000]big.Int, i, j int) bool {
	// ts[x] < ts[a] && ts[x] < ts[b]
	for x := i + 1; x < j; x++ {
		if ts[x].Cmp(&ts[i]) != -1 || ts[x].Cmp(&ts[j]) != -1 {
			return false
		}
	}
	return true
}
