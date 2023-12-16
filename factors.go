package main

import (
	"fmt"
	"math/big"
	"bufio"
	"os"
)


// factorizes as many numbers as possible into a product
// of two smaller numbers and prints the result
func print_prime_factors(number *big.Int) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// Handle the case when the number is less than or equal to 1
	if number.Cmp(one) <= 0 {
		return
	}

	// Handle the case when the number is divisible by 2
	if new(big.Int).Mod(number, two).Cmp(zero) == 0 {
		result := new(big.Int)
		result.Quo(number, two)
		fmt.Printf("%s=%s*2\n", number.String(), result.String())
		return
	}

	odd_prime := big.NewInt(3) // the starting number for odd primes
	sqrt_number := new(big.Int).Set(number)
	sqrt_number.Sqrt(sqrt_number)

	for odd_prime.Cmp(sqrt_number) <= 0 {
		if new(big.Int).Mod(number, odd_prime).Cmp(zero) == 0 {
			result := new(big.Int)
			result.Quo(number, odd_prime)
			fmt.Printf("%s=%s*%s\n", number.String(), result.String(),
						odd_prime.String())
			return
		}
		// skip this number if we go past this prime number without a match
		if odd_prime.Cmp(big.NewInt(611953)) > 0 {
			return
		}
		odd_prime.Add(odd_prime, two) // odd_prime += 2
	}

	// the number is a prime
	fmt.Printf("%s=%s*1\n", number.String(), number.String())
}

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // grabs the current line (number) in the file
		number, ok := new(big.Int).SetString(line, 0)

		// let's check if the number was converted successfully
		if ok == false {
			continue // never mind, let's get the next one
		}
		print_prime_factors(number)
	}

	// check for file reading errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
}
