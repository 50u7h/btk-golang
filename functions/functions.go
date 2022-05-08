package functions

func Topla(sayi1 int, sayi2 int) int {
	return sayi1 + sayi2
}

func Cikar(sayi1 int, sayi2 int) int {
	return sayi1 - sayi2
}

func Bol(sayi1 int, sayi2 int) float32 {
	return float32(sayi1 / sayi2)
}

func Carp(sayi1 int, sayi2 int) float32 {
	return float32(sayi1 * sayi2)
}

func DortIslem(sayi1 int, sayi2 int) (int, int, float32, int) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	bolüm := float32(sayi1 / sayi2)
	carpim := sayi1 * sayi2

	return toplam, cikarim, bolüm, carpim
}

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	return toplam
}
