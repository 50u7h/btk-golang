package pointers

import "fmt"

func Demo13(sayi *int) {
	*sayi++
	fmt.Println("Demodaki sayı: ", *sayi)
	fmt.Println("Demodaki sayının adresi: ", &sayi)
}

func Demo14(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayi: ", sayilar[0])
}
