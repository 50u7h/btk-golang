package defer_statement

import "fmt"

func A() {
	fmt.Println("A Çalıştı")
}

func B() {
	fmt.Println("B Çalıştı")
}

func C() {
	defer A()
	defer B()
	fmt.Println("C Çalıştı")
}

func Demo19(sayi int32) string {
	DeferFunc()
	if sayi%2 == 0 {
		return "çift sayı"
	}
	if sayi%2 == 1 {
		return "tek sayı"
	}

	return "belli değil"
}

func DeferFunc() {
	fmt.Println("BİTTİ")
}

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("loglandı")
}

func Demo20() {
	p := product{
		productName: "Laptop",
		unitPrice:   4599,
	}
	defer p.save()
	p = product{
		productName: "PC",
		unitPrice:   9999,
	}
	fmt.Println("İşlem başarılı")
}
