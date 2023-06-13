package main

import (
	"fmt"
)

var a int //can be used inside and outside of functions
var b int = 2
var c = 3

func main() {
	a = 1
	d := "test" //can be used only inside this function
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
