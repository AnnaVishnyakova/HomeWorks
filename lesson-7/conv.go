package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for y := <-naturals; y <= 30; y++ {
			// if !ok {
			// 	break // канал закрыт и пуст
			// }
			squares <- y * y
		}
		close(squares)
	}()

	// печать
	for {
		i, ok := <-squares
		if ok == false {
			fmt.Println("End")
			break // канал закрыт и пуст
		} else {
			fmt.Println(ok, i)

		}

		// fmt.Println(i)
		// if !ok {
		// 	fmt.Println("over")
		// 	//close(squares)
		// }

	}
}
