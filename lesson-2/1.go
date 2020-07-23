package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Первое задание
	fmt.Println("Превое задание на нечетное/четное")
	var i int
	fmt.Println("Введите число для опредления четное/нечетное:")
	fmt.Scanln(&i)

	if i%2 == 0 {

		fmt.Println(i, "Четное")

	} else {

		fmt.Println(i, "Нечетное")

	}

	//Второе задание
	fmt.Println("Второе задание на определение деления на 3")
	var g int
	fmt.Println("Введите число для опредления делится/неделится на 3:")
	fmt.Scanln(&g)
	t := strconv.Itoa(g)
	if g%3 == 0 {

		fmt.Printf("Число %v делится нацело на 3 \n", t)

	} else {

		fmt.Println("Число не делится нацело на 3")

	}
	//Третье задание
	for j, k := 0, 1; j < 100; j, k = j+k, j {
		fmt.Println(j)
	}

}
