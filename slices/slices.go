package slices

import "fmt"

func Demo7() {
	names := make([]string, 3)

	names[0] = "samet"
	names[1] = "guney"
	names[2] = "temas"
	names = append(names, "yenug")
	fmt.Println(names)
}

func Demo8() {
	cities := []string{"Trabzon", "İstanbul", "Ankara", "İzmir", "Lodz"}
	citiesCopy := make([]string, len(cities))
	copy(citiesCopy, cities)
	fmt.Println(citiesCopy)
	fmt.Println(citiesCopy[1:3])

}
