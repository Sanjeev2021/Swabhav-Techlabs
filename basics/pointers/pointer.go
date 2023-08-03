package main

import "fmt"

type Foo struct {
	bar int
}

func main() {
	var foo *Foo
	fmt.Println(foo)
	foo = new(Foo)
	fmt.Println(foo)
	fmt.Println((*foo).bar)
}
