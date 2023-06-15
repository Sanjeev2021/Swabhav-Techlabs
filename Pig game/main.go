package main

import (
	"fmt"
	"math/rand"
)

// global declare using long hand method
var turn int = 1
var totalScore int = 0


func main() {
	greet()
	for totalScore < 20 {
		fmt.Printf("\nTurn %d:\n", turn)
		fmt.Println("Welcome to the game of Pig!")
		fmt.Println("\nEnter 'r' to roll again, 'h' to hold.")

		turnManagement()
		turn++
	}

	fmt.Printf("\nYou win!! You finished in %d turns\n", turn)
	fmt.Println("Game Over")
}

// Generates a random integer between 1 and 6 (Inclusive)
func rollDice() int {
	dice := rand.Intn(6) + 1
	return dice
}

// Just greets the user with instructions and rules
func greet() {
	fmt.Println("Let's Play PIG!")
	fmt.Println("\n* See how many turns it takes you to get to 20.")
	fmt.Println("* Turn ends when you hold or roll a 1.")
	fmt.Println("* If you roll a 1, you lose all points for the turn.")
	fmt.Println("* If you hold, you save all points for the turn.")
}

// Takes and processes user input
func takeInput() string {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		// If input is not r or h
		if input != "r" && input != "h" {
			fmt.Println("Invalid input!")
			continue
		}

		break
	}
	return input
}

// Manages each turn 
func turnManagement() {
	turnScore := 0
	for {
		choice := takeInput()
		if choice == "r" {
			dice := rollDice()
			if dice == 1 {
				fmt.Println("Turn Over. No score")
				return
			}
			turnScore += dice
			showScore(turnScore)
		} else {
			break
		}
	}
	totalScore += turnScore
	fmt.Printf("Your turn score is %d and your total score is %d\n", turnScore, totalScore)
}

// Shows users their score after a successful roll
func showScore(turnScore int) {
	fmt.Printf("\nYour turn score is %d and your total score is %d \nIf you hold, you will have %d points.\nEnter 'r' to roll again, 'h' to hold.\n", turnScore, totalScore, totalScore + turnScore)
}