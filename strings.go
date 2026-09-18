package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Afakih Fajduwani", "fakih"))
	fmt.Println(strings.Split("Afakih Fajduwani", " "))
	fmt.Println(strings.ToLower("Afakih Fajduwani"))
	fmt.Println(strings.ToUpper("Afakih Fajduwani"))
	fmt.Println(strings.Trim("		Afakih		", " "))
	fmt.Println(strings.ReplaceAll("Afakih Fajduwani Afakih Fajduwani", "Afakih", "Laras"))
}
