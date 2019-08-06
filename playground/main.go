package main

import (
	"fmt"
)

func test() {
	var flt = 2.34
	var name = "Scotch"
	a, b := true, false

	fmt.Printf("%T \n", flt)
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
}

var jj = "JOe joel"

func main() {
	test()
	fmt.Println(jj)
}
