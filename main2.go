package main

import (
	"fmt"
	"time"
)

func Proceso(id int, ig chan int, c chan int) {

	i := int(0)

	for {
		//c := make(chan int)

		i = i + 1
		fmt.Println("In for")
		c <- i
		ig <- id

		fmt.Println("Sent channel")

		//time.Sleep(time.Millisecond * 500)
	}

}
func Shower(id chan int, c chan int) {
	fmt.Println("Enter 2 ")
	msg := <-c
	msg2 := <-id
	for {
		fmt.Println("In for 2")
		//msg := <-c
		fmt.Println(msg2, msg)
		time.Sleep(time.Millisecond * 500)
		msg++
	}
}

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	var id chan int = make(chan int)
	var c chan int = make(chan int)

	go Proceso(100, id, c)
	go Proceso(99, id, c)

	go Shower(id, c)

	var input2 string
	fmt.Scanln(&input2)

	/*
		var input1 string
		fmt.Scanln(&input1)

		var c chan string = make(chan string)

		go pinger(c)
		go ponger(c)
		go printer(c)

		var input string
		fmt.Scanln(&input)
	*/
}
