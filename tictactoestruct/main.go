package main

import (
	"fmt"

	game "tictactoestruct/Game"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from", r)
		}
		fmt.Println("End of the game")
	}()
	var result string
	var err error

	game1 := game.NewGame("Sanjeev", "Sahil")

	result, err = game1.Play(0)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(1)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(2)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(3)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(6)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(4)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(7)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

	result, err = game1.Play(5)
	if err != nil {
		print(err)
	}
	fmt.Println(result)

}
