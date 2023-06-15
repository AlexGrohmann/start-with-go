package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1)) //length
	fmt.Println(cap(myslice1)) //capacity
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
