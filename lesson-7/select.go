package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan string)

	go func() {

		var q string
		fmt.Scanln(&q)

		if q == "exit" {

			cancel <- q

		} else {
			fmt.Println("exit<>")
		}
		// read a single byte
		// cancel <- 10

	}()

	fmt.Println("Countdown started. Hit [Enter] to cancel")

	tick := make(<-chan time.Time)
	tick = time.Tick(3 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case moment := <-tick:
			fmt.Println(moment.Format("15:04:05"), countdown)
		case <-cancel:
			time.Sleep(time.Second * 3)
			fmt.Println("Launch canceled!")
			return
		}
	}
	fmt.Println("The rocket starts!")
}
