package main

import (
	"fmt"
	"time"
)

func mp() {

	m := make(map[string]int)
	m["1"] = 0
	m["2"] = 0
	m["3"] = 0
	m["4"] = 0

	for {
		fmt.Println("id 1", "sum", m["1"])
		m["1"] = m["1"] + 1
		fmt.Println("id 2", "sum", m["2"])
		m["2"] = m["2"] + 1
		fmt.Println("id 3", "sum", m["3"])
		m["3"] = m["3"] + 1
		fmt.Println("id 4", "sum", m["4"])
		m["4"] = m["4"] + 1
		time.Sleep(time.Second * 2)

	}

}

func main() {

	go mp()

	var input1 string
	fmt.Scanln(&input1)

	/*
		var s []int

		m := make(map[string]int)

		m["k1"] = 7
		m["k2"] = 13

		fmt.Println("map", m)

		v1 := m["k1"]
		s = append(s, v1)
		fmt.Println(s)

		fmt.Println("v1 :", v1)
		fmt.Println("map", m)

		fmt.Println("len", len(m))

			delete(m, "k2")
			fmt.Println("map:", m)

			_, prs := m["k2"]
			fmt.Println("prs:", prs)

			n := map[string]int{"foo": 1, "bar": 2}
			fmt.Println("map:", n)
	*/
}
