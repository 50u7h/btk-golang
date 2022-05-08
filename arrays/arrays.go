package arrays

import "fmt"

func Demo4() {
	var numbers [5]int
	numbers[1] = 61
	fmt.Println(numbers)
}

func Demo5() {
	cities := []string{"Trabzon", "İstanbul", "Ankara", "İzmir", "Lodz"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}

func Demo6() {
	numbers := []int{16, 72, 38, 94, 50}
	biggest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > biggest {
			biggest = numbers[i]
		}
	}
	fmt.Println("En büyük sayı: ", biggest)
}
