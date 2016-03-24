package main

import "fmt"

func main(){
	temp1 := 1; // Var to store the number in the n - 2 position
	temp2 := 2; // Var to store the number in the n - 1 position

	sum := 2 // Begin with '2' because is the firts even term in fibonacci sequence
	tempSum := 0
	limitNumber := 4000000

	for {
		tempSum = temp1 + temp2 // Calculating number n with the sum of n-2 and n-1 
		temp1 = temp2 // n - 1 is the new n - 2
		temp2 = tempSum // n is the new n - 1
		if(temp2 > limitNumber) { // Is temp2 higher that limitNumber ?
			break
		} else if ( temp2 % 2 == 0) { // Is temp2 a even number ?
			sum += tempSum
		}

//		fmt.Printf("1: %d , 2: %d sum: %d \n",temp1,temp2,sum) // Yeah... I was debugging...
	}

	fmt.Println(sum)
}
