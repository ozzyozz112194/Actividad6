package main

import (
	"fmt"
	"time"
)

var s []uint64

type pro struct {
	id  uint64
	sum uint64
}

func recive(id uint64, sum uint64) {
	e := pro{id, sum}
	s = append(s, e.id, e.sum)

}

func sum(c chan uint64) {
	var i int = 1

	for {
		s[i] = s[i] + 1
		c <- s[i]
		i += 2

		if i == len(s)+1 {

			i = 1
		}

	}
}

func show(c chan uint64) {
	var i int = 0
	var b int = 0

	for b != 15 {
		msg := <-c
		fmt.Println("id", s[i], "sum", msg)
		time.Sleep(time.Millisecond * 500)
		i += 2

		if i == len(s) {
			i = 0
		}
		b++
		fmt.Println(s)

	}
}

func eliminate(s []uint64, index int) []uint64 {
	return append(s[:index], s[index+1:]...)

}

func searchAndDestroy(id uint64) {
	for i := range s {

		if id == s[i] {
			fmt.Println("Found and destroyed")
			s = eliminate(s, i)
			s = eliminate(s, i)
			break

		}
		i++
	}
}

func main() {

	var c chan uint64 = make(chan uint64)

	var opc string
	var cont uint64 = 1
	//var countercont uint64 = 1
	recive(cont, 0)
	cont++
	go sum(c)

	for opc != "0" {

		fmt.Println("\n1)Agregar proceso\n2)Mostrar proceso\n3)Elminar procesos\n0)Salir")
		fmt.Scanln(&opc)

		switch opc {
		case "1":
			print("Case 1\n")
			recive(cont, 0)

			cont++

		case "2":
			print("Case 2\n")
			go show(c)
			var input string
			fmt.Scanln(&input)

		case "3":
			print("Case 3\n")
			var id uint64
			fmt.Scanln(&id)
			searchAndDestroy(id)
		}

	}
	println("Out")
	fmt.Println(s)
}

/*
	var c chan uint64 = make(chan uint64)
	var id uint64 = 1
	recive(1, 0)
	go sum(c)
	recive(2, 0)
	fmt.Println(s)

	//searchAndDestroy(id)
	//fmt.Println(s)
	var input string
	fmt.Scanln(&input)
	fmt.Println(s)

	fmt.Println(s)

	fmt.Println(s)

	go show(c)
	//searchAndDestroy(id)
	//fmt.Println(s)
	var input2 string
	fmt.Scanln(&input2)
	fmt.Println(s)
	searchAndDestroy(id)
	fmt.Println(s)

}
*/
