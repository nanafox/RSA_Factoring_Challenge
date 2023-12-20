package main

import (
	"fmt"
	"math/big"
)

// A hash map to store prime numbers for easy look up
var resultsCache map[string]string

func init() {
	// initialize the hash map
	resultsCache = make(map[string]string)
}

// factorizes as many numbers as possible into a product
// of two smaller numbers and prints the result
func printPrimeFactors(number *big.Int, oddPrime *big.Int) int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	limit := big.NewInt(611953)
	stopLoop := 20000000

	// now I don't care what really happens, let us go until we can't
	if oddPrime.Cmp(limit) >= 0 {
		limit.Set(big.NewInt(30597650))
	} else {
		stopLoop = 300000000
	}

	// Handle the case when the number is less than or equal to 1
	if number.Cmp(one) <= 0 {
		return 0
	}

	// check for cached values
	cachedResult, found := resultsCache[number.String()]
	if found {
		fmt.Printf("%s=%s\n", number.String(), cachedResult)
		return 0
	}

	// Handle the case when the number is divisible by 2
	if new(big.Int).Mod(number, two).Cmp(zero) == 0 {
		quotient := new(big.Int)
		quotient.Quo(number, two)
		resultsCache[number.String()] = quotient.String() + "*2"

		// print the result and exit from here
		fmt.Printf("%s=%s*2\n", number.String(), quotient.String())
		return 0
	}

	sqrtNumber := new(big.Int).Set(number)
	sqrtNumber.Sqrt(sqrtNumber)

	loopCounter := 0
	for oddPrime.Cmp(sqrtNumber) <= 0 {
		if new(big.Int).Mod(number, oddPrime).Cmp(zero) == 0 {
			quotient := new(big.Int)

			// this is an expensive operation so we'd save the result for later
			quotient.Quo(number, oddPrime)

			// let's save the result for later
			resultsCache[number.String()] = quotient.String() + "*" +
				oddPrime.String()
			fmt.Printf("%s=%s*%s\n", number.String(), quotient.String(),
				oddPrime.String())
			return 0
		}
		// skip this number if we go past this prime number without a match
		if oddPrime.Cmp(limit) > 0 || loopCounter > stopLoop {
			// fmt.Printf("Limit reached for [%s]\n", number.String())
			return 1
		}
		oddPrime.Add(oddPrime, two) // oddPrime += 2
		loopCounter++
	}

	// the number is a prime
	fmt.Printf("%s=%s*1\n", number.String(), number.String())
	return 0
}
