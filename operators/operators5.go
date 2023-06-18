package main

import (
	"fmt"
)

func main() {
	var x = 5
	var y = 3
	fmt.Println(x > y) // returns 1 (true) because 5 is greater than 3
}

// ==	Equal to	x == y
// !=	Not equal	x != y
// >	Greater than	x > y
// <	Less than	x < y
// >=	Greater than or equal to	x >= y
// <=	Less than or equal to	x <= y

//Logical operators

// && 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10
// || 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4
// !	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)
