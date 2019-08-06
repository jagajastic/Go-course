package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	license bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `Says, "Shaken, not stirred"`)
}

func main() {
	sap := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sap.speak()

}
