package visibility

import (
	"math/big"
	"testing"
)

func TestHorizontal(t *testing.T) {
	// Input
	arr := [10000]big.Int{
		*big.NewInt(2),
		*big.NewInt(4),
		*big.NewInt(2),
		*big.NewInt(4),
		*big.NewInt(6),
		*big.NewInt(2),
	}

	visibles := [][2]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			vis := Horizontal(arr, i, j)
			if vis {
				visibles = append(visibles, [2]int{i + 1, j + 1})
			}
		}
	}
	// Result must be
	// {1 -> 2, 2 -> 3, 2 -> 4, 3 -> 4, 4 -> 5, 5 -> 6}
	if !(visibles[0] == ([2]int{1, 2}) && visibles[1] == ([2]int{2, 3}) &&
		visibles[2] == ([2]int{2, 4}) && visibles[3] == ([2]int{3, 4}) &&
		visibles[4] == ([2]int{4, 5}) && visibles[5] == ([2]int{5, 6})) {
		t.Error("wrong value")
	}
	t.Log(visibles)

	// Second test
	arr = [10000]big.Int{
		*big.NewInt(2),
		*big.NewInt(6),
		*big.NewInt(4),
		*big.NewInt(2),
		*big.NewInt(4),
		*big.NewInt(6),
	}

	visibles = [][2]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			vis := Horizontal(arr, i, j)
			if vis {
				visibles = append(visibles, [2]int{i + 1, j + 1})
			}
		}
	}
	// Result must be
	// {1 -> 2, 2 -> 3, 2 -> 6, 3 -> 4, 3 -> 5, 4 -> 5, 5 -> 6}
	if !(visibles[0] == ([2]int{1, 2}) && visibles[1] == ([2]int{2, 3}) &&
		visibles[2] == ([2]int{2, 6}) && visibles[3] == ([2]int{3, 4}) &&
		visibles[4] == ([2]int{3, 5}) && visibles[5] == ([2]int{4, 5}) &&
		visibles[6] == ([2]int{5, 6})) {
		t.Error("wrong value")
	}
	t.Log(visibles)

}
