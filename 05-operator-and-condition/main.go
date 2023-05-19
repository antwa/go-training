package main

import (
	"fmt"
)

func main() {
	// arithmetic operatior
	// +,-,*,/,%,++,--
	var a, b, c, d, e, f = 1, 2, 0, 0, 0, 0

	c = a + b
	fmt.Println(c)

	d = a - b
	fmt.Println(d)

	e = a * b
	fmt.Println(e)

	f = a % b
	fmt.Println(f)

	// condition
	// >,>=,<,<=,==,!=
	// logical operator
	// &&,||,|

	if a > 1 {
		fmt.Println("nilai a lebih besar dari 1")
	} else {
		fmt.Println("nilai a lebih kecil dari 1")
	}

}
