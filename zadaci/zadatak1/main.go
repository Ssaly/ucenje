package main

import (
	"fmt"
)

func main() {

	fmt.Println("Zadatak 1 ucenje")

	for i := 2000; i <= 3200; i++ {

		if i%7 == 0 && i%5 != 0 {
			fmt.Printf(" %v, ", i)
		}

	}

}
