package main

import (
	"os"
)

func main() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	panic("a problem")

	_, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
}
