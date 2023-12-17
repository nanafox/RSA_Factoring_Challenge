package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// A hash map to store prime numbers for easy look up
var result_cache map[string]string

// reads from a file and calls another function perform and print
// prime factors
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "USAGE: factors file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Can't open file %s\n", os.Args[1])
		os.Exit(1)
	}
	defer file.Close()

	// initialize the hash map
	result_cache = make(map[string]string)

	var big_primes = []*big.Int{}
	size_of_primes := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // grabs the current line (number) in the file
		number, ok := new(big.Int).SetString(line, 0)

		// let's check if the number was converted successfully
		if !ok {
			continue // never mind, let's get the next one
		}

		val := print_prime_factors(number, big.NewInt(3))

		// keep numbers with large with prime factors for later
		if val == 1 {
			big_primes = append(big_primes, number)
			size_of_primes++
		}
	}

	fmt.Printf("Size of big primes: [%d]\n", size_of_primes)
	/* work on the numbers with very big prime factors */
	for i := 0; i < size_of_primes; i++ {
		print_prime_factors(big_primes[i], big.NewInt(611953))
	}

	// check for file reading errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
}
