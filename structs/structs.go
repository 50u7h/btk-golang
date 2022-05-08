package structs

import "fmt"

func Demo15() {
	fmt.Println(product{"Bilgisayar", 10999, "Monster"})
	fmt.Println(product{"Bilgisayar", 12999, ""})
	fmt.Println(product{name: "Telefon", unitPrice: 5699})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}

type customer struct {
	name        string
	mail        string
	phoneNumber string
}

func (c customer) add() {
	fmt.Println("Müşteri eklendi.")
	fmt.Println("İsim: ", c.name, "\nMail: ", c.mail, "\nTel : ", c.phoneNumber)
}

func Demo16() {
	c := customer{
		name:        "Samet Güney",
		mail:        "samet@guney.com",
		phoneNumber: "+90534961171",
	}
	c.add()
}
