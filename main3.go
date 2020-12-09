package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go f(0)
	go f(1)
	go f(2)
	go f(3)

	var input string
	fmt.Scanln(&input)
}
