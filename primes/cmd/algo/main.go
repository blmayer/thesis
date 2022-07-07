package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"primes/internal/visibility"
)

const usage = `Visibility algorithms utility
Usage: algo algorithm treads file.csv number

Meaning of arguments:
 algorithm:	algorithm to run
 threads:	number of threads to use
 file:		path to csv file containing primes
 number:	how many primes from the file to consider

Options for algorithm:
 natural
 horizontal

Outputs:
 This utility, upon successful run, writes to terminal

Example:
 algo natural 50 11-primes.csv 1000
`

var finished chan int

func main() {
	// Read argument
	if len(os.Args) < 5 {
		fmt.Println("missing argument:")
		fmt.Println(usage)
		return
	}

	// k, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	log.Fatal("error reading threads argument:", err)
	// }

	n, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatal("error reading number argument:", err)
	}

	primes := make([]big.Int, n)
	file, err := os.Open(os.Args[3])
	if err != nil {
		log.Fatal("error opening primes file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		// Add number to list
		_, ok := primes[counter].SetString(scanner.Text(), 10)
		if !ok {
			fmt.Println("failed to parse")
		}
		counter++
		if counter == n {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate differences
	diffs := [10000]big.Int{}
	for counter = 0; counter < n-1; counter++ {
		temp := big.NewInt(0)
		diffs[counter] = *temp.Sub(&primes[counter+1], &primes[counter])
	}

	// Select algorithm to run
	finished = make(chan int)
	switch os.Args[1] {
	case "natural":
		// Calculate natural visibility
		for i := 0; i < n-1; i++ {
			go runNatural(&diffs, n, i)
		}

	case "horizontal":
		// Calculate horizontal visibility
		for i := 0; i < n-1; i++ {
			go runHorizontal(&diffs, n, i)
		}

	default:
		fmt.Fprintln(os.Stderr, "algorithm not found")
	}

	// Wait all goroutines
	counter = 0
	for counter < n-1 {
		counter += <-finished
	}
}

func runNatural(diffs *[10000]big.Int, n, i int) {
	for j := i + 1; j < n-1; j++ {
		vis := visibility.Natural(*diffs, i, j)
		if vis {
			fmt.Printf("%d,%d\n", i+1, j+1)
		}
	}
	finished <- 1
}

func runHorizontal(diffs *[10000]big.Int, n, i int) {
	for j := i + 1; j < n-1; j++ {
		vis := visibility.Horizontal(*diffs, i, j)
		if vis {
			fmt.Printf("%d,%d\n", i+1, j+1)
		}
	}
	finished <- 1
}
