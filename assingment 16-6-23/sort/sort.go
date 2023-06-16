package sort

import (
	"fmt"
	"sort"
)

func main(){
	numbers := []int{3,5,6,7,1,2,4,5,9}
	sortArray(numbers)
	fmt.Println("The sorted numbers are :", numbers)
}

func sortArray(arr []int){
	sort.Ints(arr)
}