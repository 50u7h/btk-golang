package maps

import "fmt"

func Demo9() {
	sozluk := make(map[string]string)

	sozluk["book"] = "masa"
	sozluk["door"] = "kapı"
	sozluk["wall"] = "duvar"

	fmt.Println(sozluk["wall"])

	deger, varMi := sozluk["wall"]
	fmt.Println(deger, varMi)

	sozluk2 := map[string]string{"mouse": "fare", "keyboard": "klavye"}
	fmt.Println(sozluk2)
}
