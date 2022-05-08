package conditionals

import "fmt"

func Workshop1() {
	var x, y, z int = 3, 2, 1
	var big int = x

	if y > big {
		big = y
	}
	if z > big {
		big = z
	}

	fmt.Printf("En büyük sayı: %v", big)
}
