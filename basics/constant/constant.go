package main

import "fmt"

// const (
// 	User    string = "Admin"
// 	Product string = "Product"
// )

//const incr = iota

const (
	i = iota
	j = iota
	k = iota
)

const (
	a = iota // 0
	b        // 1
	c        // 2
	d        // 3
)

func main() {
	// const i int = 12

	// const j float32 = 3.14
	// const k string = "sanjeev"
	// const l bool = true
	// var a int = 12
	// fmt.Println(i + a)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
