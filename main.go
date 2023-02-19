package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := generateSecretNum()

	recursiveGuessing(secret)
}

func recursiveGuessing(secret int) {
	// user guess
	input := getUserInput()

	// comparison
	matching := checkMatching(secret, input)

	if !matching {
		fmt.Println("Your Guess is Wrong.")

		recursiveGuessing(secret)

	} else {
		fmt.Println("Your Guess is Correct. The result is: ", secret)
	}
}

func checkMatching(secret int, input int) bool {

	return secret == input
}

func getUserInput() int {
	fmt.Print("Please type your guess here: ")

	var input int

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Your Input Failed. Try again")
	} else {
		fmt.Println("Guess Done: ", input)
	}

	return input

}

func generateSecretNum() int {
	rand.Seed(time.Now().Unix()) // add rand time

	return rand.Int() % 10
}
