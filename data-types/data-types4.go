package main

import (
	"fmt"
)

func main() {
	var x uint = 500  //only positive values
	var y uint = 4500 //only positive values
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}

// uint	Depends on platform:
// 32 bits in 32 bit systems and
// 64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
// 0 to 18446744073709551615 in 64 bit systems
// uint8	8 bits/1 byte	0 to 255
// uint16	16 bits/2 byte	0 to 65535
// uint32	32 bits/4 byte	0 to 4294967295
// uint64	64 bits/8 byte	0 to 18446744073709551615
