package loops

import "fmt"

func Workshop2() {
	sayi := 0
	fmt.Print("Bir sayı giriniz: ")
	fmt.Scan(&sayi)
	var asalmi bool = true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalmi = false
		}
	}

	if asalmi {
		fmt.Print("Bu bir asal sayıdır")
	} else {
		fmt.Print("Asal sayı değil !!!")
	}

}
