package main

import (
	"fmt"
	"math/big"
	"os"
	"primes/internal/generator"
	"strconv"
)

const usage = `d-prime generator utility
usage: generator d n

meaning os parameters:
 d:	number of divisors
 n: how many primes are to be generated

example:
 generator 2 10000 > 2-primes.txt

to generate 10000 2-primes and write them to 2-primes.txt file
`

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "missing argument")
		fmt.Fprintln(os.Stderr, usage)
		return
	}
	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading divisor: "+err.Error())
		return
	}

	lim, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading upper limit:"+err.Error())
		return
	}

	i := big.NewInt(1)
	one := big.NewInt(1)
	for lim > 0 {
		divisors := generator.NumOfDivisors(i)
		if divisors == d {
			fmt.Println(i.String())
			lim--
		}
		i.Add(i, one)
	}
}
