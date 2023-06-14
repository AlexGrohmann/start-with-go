package main

import (
	"fmt"
)

func main() {
	var x float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x)
}

// float32	32 bits	-3.4e+38 to 3.4e+38.
// float64	64 bits	-1.7e+308 to +1.7e+308.
