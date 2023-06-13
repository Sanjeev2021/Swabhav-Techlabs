package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <=1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i:=2 ; i<=limit ; i++{
		if (n%i ==0){
			return false;
		}
	}
	return true

}

func main() {

 number:= 21

 if isPrime(number){
	fmt.Println(number , "it is prime")
 } else {
	fmt.Println(number , "it is not prime")
 }
}