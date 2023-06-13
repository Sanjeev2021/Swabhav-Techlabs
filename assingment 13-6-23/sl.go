package main

import (
	"fmt"
	"sort"
)

func main()  {
	arr:= []int{1,2,3,4,5,6}
	sort.Ints(arr)
	if len(arr)<=1 {
		fmt.Println("no second largest number")
	}else {
		fmt.Println("second largest number is", arr[len(arr)-2])
	}

}