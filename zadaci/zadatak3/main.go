package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter number: ")
	var n int
	fmt.Scan(&n)

	mapa := make(map[int]int)

	for i := 1; i <= n; i++ {
		mapa[i] = i * i
	}
	fmt.Println(mapa)

}
