package main

import (
	"fmt"
	"strconv" // using new library...
)

func main(){

	var begin int = 900
	var limit int = 999
	var mayorPalindrome int = 0

// Brute Force Solution
	for i := begin ; i <= limit ; i++ {
		for j := i ; j <= limit ; j++ {
			product := i*j
			if isPalindromic(product) {
				if i*j > mayorPalindrome {
					mayorPalindrome = product
				}
			}
		}
	}

	fmt.Println(mayorPalindrome)
}

func isPalindromic(number int) bool {
	var stringNumber string = strconv.Itoa(number) // Convert int to string
	var isPalindromic bool = true

	for i := 0 ; i <= (len(stringNumber) - 1)/2 ; i++ {
		if stringNumber[i] != stringNumber[len(stringNumber)-1-i] {
			isPalindromic = false
			break
		}
	}
	return isPalindromic
}
