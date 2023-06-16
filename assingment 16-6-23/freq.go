package main

import "fmt"

func main() {
	number := []int{2,3,4,5,6,7,2,4,6}
	target := 5

	count := countOccurrences(number, target)
	fmt.Println("The number of occurrences of target in the array is: ", target, count)
}

func countOccurrences(numbers []int, target int) int {
	count := 0
	for _, num := range numbers {
		if num == target {
			count++
		}
	}
	return count
}