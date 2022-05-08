package loops

import "fmt"

func Workshop3() {
	sayi1 := 0
	sayi2 := 0
	sonuc1 := 0
	sonuc2 := 0

	fmt.Print("İki sayı giriniz: ")
	fmt.Scan(&sayi1, &sayi2)

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			sonuc1 += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			sonuc2 += i
		}
	}

	if sayi1 == sonuc2 && sayi2 == sonuc1 {
		fmt.Println("Bunlar arkadaş sayıdır")
	} else {
		fmt.Println("Değildir")
	}
}
