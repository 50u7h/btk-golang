package conditionals

import "fmt"

func Demo2() {
	var balance float32 = 1000
	var withdraw float32
	var kalan float32

	print("Çekmek istediğiniz tutar: ")
	fmt.Scanf("%g", &withdraw)

	kalan = balance - withdraw

	if kalan < 0 {
		fmt.Println("Bakiye yetersiz!!!")
	} else {
		fmt.Printf("Kalan bakiye: %g", kalan)
	}

}
