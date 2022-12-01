package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter number: ")
	var broj int
	fmt.Scan(&broj)

	rez := computeFactorial(broj)

	fmt.Printf("Factorial broja %v je %v ", broj, rez)
}

func computeFactorial(broj int) int {
	factorial := 1
	for i := 1; i <= broj; i++ {
		factorial *= i
	}
	return factorial
}
