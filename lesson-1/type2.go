package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Println("Введите сторону а:")
	fmt.Scanln(&a)
	fmt.Println("Введите сторону b:")
	fmt.Scanln(&b)
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println("Катет равен:", c)
	s := (a * b) / 2
	fmt.Println("Площадь равна:", s)
	p := a + b + c
	fmt.Println("Периметр равен:", p)

}
