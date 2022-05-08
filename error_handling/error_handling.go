package error_handling

import (
	"errors"
	"fmt"
	"os"
)

func Demo21() {
	f, err := os.Open("demo.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadığğğ:", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadığğğ")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}

func TahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı girin")
	}
	return "Bildiniz", nil
}

func Demo23() {
	fmt.Println(TahminEt(50))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-10))
}

type BorderException struct {
	message   string
	parametre int
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parametre, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &BorderException{
			message:   "sınırların dışında kaldı",
			parametre: tahmin,
		}
	}
	return "bildiniz", nil
}
