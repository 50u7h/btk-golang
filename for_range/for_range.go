package for_range

import "fmt"

func Demo10() {
	sehirler := []string{"Trabzon", "London", "Lodz"}

	for i, s := range sehirler {
		fmt.Println((i + 1), s)
	}
}

func Demo11() {
	toplam := 0
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i2 := range sayilar {
		if i2%2 == 1 {
			toplam += i2
		}
	}
	fmt.Println("Tek sayıların toplamı: ", toplam)
}

func Demo12() {
	sozluk := map[string]string{"mouse": "fare", "keyboard": "klavye", "table": "masa", "chair": "sandalye"}

	for s1, s2 := range sozluk {
		fmt.Println(s1, "=", s2)
	}
}
