package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func main() {
	var a [15]int
	for i := 0; i < 15; i++ {
		fmt.Printf("Element: %d %dd\n", i, a[i])
	}

	// slice holds list of things like array in js
	xi := []int{2, 3, 4, 5, 6}
	for i := 0; i < len(xi); i++ {
		fmt.Println("count", xi[i])
	}
	fmt.Println(xi)

	// map is a key value data type
	m := map[string]string{
		"name":    "Tood",
		"Surname": "Mcloed",
	}

	fmt.Println(m)

	p1 := person{
		"James",
		"Kalouw",
	}

	fmt.Println(p1)
}
