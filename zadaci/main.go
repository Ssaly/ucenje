package main

import (
	"fmt"
)

func main() {

	fmt.Println("Exercise 001")

	res, _ := suma(5, 5)

	fmt.Println(res)

}

func suma(a, b int) (int, error) {

	return a + b, nil

}
