package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `Says, Good morning, James`)
}

func main() {
	p1 := person{
		"Miss",
		"Kelleb",
	}

	p1.speak()

}
