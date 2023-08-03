package main

import "fmt"

func main() {
	// i, j := 0, 0
	// for i < 10 {
	// 	fmt.Println(i, j)
	// 	i, j = i+1, j+1
	// 	if i == 6 && j == 6 {
	// 		continue
	// 	}
	// }
	// fmt.Println(i, j)
	// i,j = i+1 , j +1
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Println(i * j)
	// 	}
	// }

	a := []int{1, 2, 3, 4, 5, 6, 7}
	for k, v := range a {
		fmt.Println(k, v)
	}

}
