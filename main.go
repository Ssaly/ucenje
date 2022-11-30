package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {

		fmt.Println("Welcome to the game! ")

	} else {
		fmt.Println("You cannot play!\n")
		return
	}

	score := 0
	num_que := 2
	fmt.Printf("Which club has won the most Champions League titles? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)
	if answer+" "+answer2 == "Real Madrid" {
		fmt.Println("Correct!")
		score += 1
	} else if answer+" "+answer2 == "real madrid" {
		fmt.Println("Correct!")
		score += 1
	} else if answer+" "+answer2 == "Real madrid" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Which outfield player appeared in the Champions League final in three different decades? ")
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "Ryan Giggs" {
		fmt.Println("Correct!")
		score += 1
	} else if answer+" "+answer2 == "ryan giggs" {
		fmt.Println("Correct!")
		score += 1
	} else if answer+" "+answer2 == "ryan gigs" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")

	}

	fmt.Printf("You scored %v out of %v.", score, num_que)
	percent := (float64(score) / float64(num_que)) * 100

	fmt.Printf("You scored: %v%%.", percent)
}
