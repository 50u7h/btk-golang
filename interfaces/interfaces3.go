package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo22(x interface{}) {
	dogrula(x)
}
