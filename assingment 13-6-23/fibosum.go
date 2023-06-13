package main 

import "fmt"

func fibonacciSum(number int ) int{
	if number <= 0 {
		fmt.Println("0")
		} 
		num1 , num2 := 0,1;
		sum := num1 + num2

		for i := 2; i < number; i++ {
			num3 := num1 + num2;
			sum += num3

			
			num1 = num2
			num2 = num3
		}

		return sum

}

func main() {
	n := 10
	fmt.Println("Sum of Fibonacci series ", n, fibonacciSum(n))
}