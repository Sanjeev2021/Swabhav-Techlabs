package main

import "fmt"

func main() {
	var i interface{} = 5
	switch i.(type) {
	case int:
		fmt.Println("THIS IS INT")

	case float64:
		fmt.Println("this is float type")
	case string:
		fmt.Println("this is string type")
	}
}
