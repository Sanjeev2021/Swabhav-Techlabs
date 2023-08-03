package main

import "fmt"

type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	student1 := Student{
		name:   "Sanjeev",
		rollNo: 5,
		subjects: []string{
			"Maths",
			"Physics",
			"Chemistry",
		},
	}
	student1.rollNo = 6
	fmt.Println(student1.name)
	fmt.Println(student1.subjects)
	fmt.Println(student1.rollNo)
	student2 := &student1
	student2.name = "sahil"
	fmt.Println(student1)
	fmt.Println(student2)
}
