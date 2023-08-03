package main

import (
	"fmt"
	"strconv"
)

func main() {
	//IDENT error means indention error
	//var i int
	//i = 12
	var i int = 12
	var j float32 = float32(i) // type conversion is happening over here
	k := 3.4
	// use strconv when converting from integer to string
	var foo string = strconv.Itoa(i)
	var bar bool = true
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(foo)
	fmt.Println(bar)
}
