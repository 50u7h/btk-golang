package main

import (
	"fmt"
	"golesson/project"
)

func main() {

	/*variables.Demo1()
	conditionals.Demo2()
	loops.Demo3()
	arrays.Demo4()
	arrays.Demo5()
	arrays.Demo6()
	slices.Demo7()
	slices.Demo8()*/

	/*fmt.Println(functions.Topla(1, 2))
	fmt.Println(functions.Cikar(1, 2))
	fmt.Println(functions.Bol(10, 3))
	fmt.Println(functions.Carp(4, 5))

	sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(12, 6)

	fmt.Println("Toplam: ", sonuc1)
	fmt.Println("Çıkarım: ", sonuc2)
	fmt.Println("Bölüm: ", sonuc3)
	fmt.Println("Çarpım: ", sonuc4)

	fmt.Println(functions.ToplaVariadic(1, 2, 3, 4, 5, 4, 3, 2, 1))*/

	//maps.Demo9()
	//for_range.Demo10()
	//for_range.Demo11()
	//for_range.Demo12()

	/*sayi := 20
	pointers.Demo13(&sayi)
	fmt.Println("Maindeki sayı: ", sayi)
	fmt.Println("Maindeki sayının adresi: ", &sayi)*/

	/* := []int{1, 2, 3, 4}
	pointers.Demo14(sayilar)
	fmt.Println("Maindeki sayi: ", sayilar[0])*/

	//structs.Demo15()
	//structs.Demo16()

	/*go go_routines.CiftSayilar()
	go go_routines.TekSayilar()
	time.Sleep(6 * time.Second)*/

	/*ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)

	go go_channels.CiftSayilar(ciftSayiCn)
	go go_channels.TekSayilar(tekSayiCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	toplam := ciftSayiToplam + tekSayiToplam
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Printf("Toplam => %v + %v = %v\n", ciftSayiToplam, tekSayiToplam, toplam)
	fmt.Printf("Çarpım => %v x %v = %v ", ciftSayiToplam, tekSayiToplam, carpim)*/

	//interfaces.Demo17()
	//interfaces.Demo18()

	//defer_statement.C()
	//fmt.Println(defer_statement.Demo19(10))
	//defer_statement.Demo20()

	//error_handling.Demo21()
	//interfaces.Demo22("selam")
	//error_handling.Demo23()
	//fmt.Println(error_handling.TahminEt2(101))

	//string_functions.Demo24()

	//restful.Demo25()
	//restful.Demo26()

	//project.GetAllProducts()
	project.AddProduct()

	products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
