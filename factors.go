package main

import (
	"fmt"
	"math/big"
	"bufio"
	"os"
)

var result_cache map[string]string

// factorizes as many numbers as possible into a product
// of two smaller numbers and prints the result
func print_prime_factors(number *big.Int, odd_prime *big.Int) int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	limit := big.NewInt(611953)

	// now I don't care what really happens, let us go until we can't
	if odd_prime.Cmp(limit) == 0 {
		limit.Set(big.NewInt(1000000000))
	}

	// Handle the case when the number is less than or equal to 1
	if number.Cmp(one) <= 0 {
		return 0
	}

	// check for cached values
	cached_result, found := result_cache[number.String()]
	if found {
		fmt.Printf("%s=%s\n", number.String(), cached_result)
		return 0
	}

	// Handle the case when the number is divisible by 2
	if new(big.Int).Mod(number, two).Cmp(zero) == 0 {
		quotient := new(big.Int)
		quotient.Quo(number, two)
		result_cache[number.String()] = quotient.String() + "*" +
										two.String()
		fmt.Printf("%s=%s*2\n", number.String(), quotient.String())
		return 0
	}

	sqrt_number := new(big.Int).Set(number)
	sqrt_number.Sqrt(sqrt_number)

	for odd_prime.Cmp(sqrt_number) <= 0 {
		if new(big.Int).Mod(number, odd_prime).Cmp(zero) == 0 {
			quotient := new(big.Int)

			// this is an expensive operation so we'd save the result for later
			quotient.Quo(number, odd_prime)

			// let's save the result for later
			result_cache[number.String()] = quotient.String() + "*" +
											odd_prime.String()
			fmt.Printf("%s=%s*%s\n", number.String(), quotient.String(),
						odd_prime.String())
			return 0
		}
		// skip this number if we go past this prime number without a match
		if odd_prime.Cmp(limit) > 0 {
			return 1
		}
		odd_prime.Add(odd_prime, two) // odd_prime += 2
	}

	// the number is a prime
	fmt.Printf("%s=%s*1\n", number.String(), number.String())
	return 0
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

	// initialize the hash map
	result_cache = make(map[string]string)

	var x = []*big.Int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // grabs the current line (number) in the file
		number, ok := new(big.Int).SetString(line, 0)

		// let's check if the number was converted successfully
		if ok == false {
			continue // never mind, let's get the next one
		}

		val := print_prime_factors(number, big.NewInt(3))

		// keep numbers with large with prime factors for later
		if (val == 1) {
			x = append(x, number)
		}
	}

	/* work on the numbers with very big prime factors */
	size := len(x)

	for i := 0; i < size; i++ {
		print_prime_factors(x[i], big.NewInt(611953))
	}

	// check for file reading errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
}
