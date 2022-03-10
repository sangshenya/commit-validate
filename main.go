package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]string, 10)
	for i := 0; i < 1000; i++ {
		s := strconv.Itoa(i)
		b = append(b, s)
	}
	fmt.Println(b[0])
	// fbasfjqwlhefqwlefhbqhjewfgbejsfabgdhbgsdhffasfgbasdhgalsjfhajklsfhaqjskfghasjldghsdfbhsghkasdgksdghgasghdahskfgqkgqyhghka
	//fmt.Println("fbasfjqwlhefqwlefhbqhjewfgbejsfabgdhbgsdhfgaskhgshagshdghsfghasfjaksgk", b[0])
}
