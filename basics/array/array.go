package main

import "fmt"

func main() {
	// var amounts [3]int = [3]int{10, 20, 30}
	// amt := [...]int{30, 40, 50, 56, 35, 67}
	// fmt.Printf("Amount: %v", amounts)
	// fmt.Printf("Amount: %v", amt)
	// fmt.Printf("%v\n", len(amt))
	// var amounts [3]int = [3]int{10, 20, 30}
	// a := &amounts
	// fmt.Printf("Amount: %v\n", amounts)
	// fmt.Printf("A: %v\n", a)
	// amounts[0] = 51

	// fmt.Printf("Amount: %v\n", amounts)
	// fmt.Printf("A: %v\n", a)
	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// // this method a[:] it will eliminate the element in array
	// b := a[:]
	// c := a[2:]
	// d := a[:3]
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// var identityMatrix [3][3]int = [3][3]int{
	// 	[3]int{1, 0, 0},
	// 	[3]int{0, 1, 0},
	// 	[3]int{0, 0, 1},
	// }
	// identityMatrix[1][2] = 7
	// fmt.Println(identityMatrix)
	// var a [3]int = [3]int{1, 2, 3}
	// fmt.Println(a)

	// var a []int = []int{1, 2, 3}
	// var b []int = a

	// fmt.Println(len(a))
	// fmt.Println(b)

	// make is used to make a slice
	// var a []int = make([]int, 3)

	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	var a []int = []int{1, 2, 3}
	var b []int = append(a, 5)
	fmt.Println(a)
	fmt.Println(b)
}
