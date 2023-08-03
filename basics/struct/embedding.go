package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryType     string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	computer := Computer{}
	computer.price = 50000
	computer.processorName = "Intel i5"
	computer.cores = 6
	computer.memoryCapacity = 8
	computer.memoryType = "DDR4"

	computer1 := Computer{
		Processor: Processor{
			processorName: "intel i7",
			cores:         8,
		},
		Memory: Memory{
			memoryCapacity: 16,
			memoryType:     "DDR4",
		},
		price: 20000,
	}
	fmt.Println(computer)
	fmt.Println(computer1)
}
