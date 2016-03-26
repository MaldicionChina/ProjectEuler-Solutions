package main

import "fmt"

func main(){
	number := 600851475143
//	number := 13195
	fmt.Println(primeFactors(number))
}

func primeFactors(number int) []int {
	// I found that list of primer number below to a 10000 
	// is enough to find the prime factors of 600851475143
	//.. yep.. I try several times to that short list...

	listOfPrimes := primeGenerator(10000)
	quotient := number // First quotient... the number to find prime factor

	var listPrimeFactors []int

	// When the quotient of the division is one... there are not more to divide...
	for quotient != 1 { // Is 
		// I take the number and divide the quotient for the list of prime factors until 
		// rest is equal zero
		for _,prime := range listOfPrimes {
			if quotient % prime  == 0 {
				quotient = quotient/prime // Update quotient
//				fmt.Println(prime)
				listPrimeFactors = append(listPrimeFactors,prime)
			}
		}
	}

	return listPrimeFactors
}

// Brute Force Prime Finder
// Custom function to generate prime numbers... don't expect serious implementation...
func primeGenerator(limit int) []int {
//	fmt.Println("Calculating Primes")
	listOfPrimes := []int{2}

	// This for sentence go over all odd numbers
	for number := 3; number <= limit; number += 2 {
		isPrime := true
		for divisor := 3; divisor < number; divisor++ {
			if number%divisor == 0 { // 
				isPrime = false
				break // This number isn't a prime
			}
		}

		if isPrime {
			listOfPrimes = append(listOfPrimes,number) // Add the primer number to list
		}
	}
//	fmt.Println("Finish")
	return listOfPrimes
}
