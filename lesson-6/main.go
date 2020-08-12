package main

import (
	"HomeWorks/lesson-6/Average"
	"fmt"
)

func main() {
	var inputSlise []float64
	var input float64 = 1
	for input != 0 {
		fmt.Scanln(&input)
		inputSlise = append(inputSlise, input)

	}
	res := Average.Average(inputSlise)
	fmt.Println(res)
}
