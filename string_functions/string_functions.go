package string_functions

import (
	"fmt"
	s "strings"
)

func Demo24() {
	name := "Samet"
	fmt.Println(s.Count(name, "a"))
	fmt.Println(s.Contains(name, "a"))
	fmt.Println(s.Index(name, "t"))

	fmt.Println(s.HasPrefix(name, "Sam"))
	fmt.Println(s.HasSuffix(name, "eta"))
	fmt.Println(s.Index(name, "et"))

	harfler := []string{"s", "a", "m", "e", "t"}
	sonuc := s.Join(harfler, "$")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "$", "|", 3))
	fmt.Println(s.Split(sonuc, "$"))
	fmt.Println(s.Repeat(sonuc, 5))
}
