package loops

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Workshop1() {
	x := rand.Intn(100)
	sayi := 0
	deneme := 0

	for true {
		fmt.Print("Sayıyı giriniz: ")
		fmt.Scan(&sayi)

		if sayi == x {
			fmt.Printf("Doğru!!! %v. tahminde bildin", deneme)
			break
		}
		if sayi > x {
			fmt.Println("Aşağı ")
			deneme++
		}
		if sayi < x {
			fmt.Println("Yukarı")
			deneme++
		}
	}
}
