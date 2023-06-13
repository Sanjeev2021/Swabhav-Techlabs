package main

import "fmt"

func countEvenOddZero(numbers[] int) (int , int ,int ){
	evenCount , oddCount , zeroCount := 0, 0, 0

	for i :=0 ; i<len(numbers) ; i++{
	num := numbers[i]

	switch {
	case num ==0:
		zeroCount++;
	case num%2 ==0:
		evenCount++;
	case num%2 !=0:
		oddCount++;
	}
}
 return evenCount, oddCount, zeroCount;
}

func main() {
	numbers := []int{2,4,5,6,2,0,0,18,36,46,21}

	evenCount , oddCount , zeroCount := countEvenOddZero(numbers)

	fmt.Println("even number are: \n", evenCount)
	fmt.Println("odd number are \n", oddCount)
	fmt.Println("zero number are", zeroCount)
}

//lots of errors noticed 

