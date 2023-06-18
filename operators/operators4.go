package main

import (
	"fmt"
)

func main() {
	var x = 10
	x += 5
	fmt.Println(x)
}

// =	x = 5	x = 5
// +=	x += 3	x = x + 3
// -=	x -= 3	x = x - 3
// *=	x *= 3	x = x * 3
// /=	x /= 3	x = x / 3
// %=	x %= 3	x = x % 3
// &=	x &= 3	x = x & 3
// |=	x |= 3	x = x | 3
// ^=	x ^= 3	x = x ^ 3
// >>=	x >>= 3	x = x >> 3
// <<=	x <<= 3	x = x << 3
