package main

import "fmt"

func main(){
	var totalCounting int = 1000
	var sum int = 0

	for number := 1; number < totalCounting; number++ {
		if number%3 == 0 { // is number multiply of 3 ? 
			sum += number
		//	fmt.Printf("%d ,", number) // To analize the progression..
		} else if number%5 == 0 { // is number multiply of 5 ?
			sum += number
		//	fmt.Printf("%d ,", number) // To analize the progression..
		}
	}
	fmt.Println(sum)
}
