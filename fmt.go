package main

import "fmt"

func main() {
	firstName := "Afakih"
	lastName := "Fajduwani"

	fmt.Println("Hello'", firstName, lastName, "'")    // ada spasi
	fmt.Printf("Hello '%s %s'\n", firstName, lastName) // tidak ada spasi
}
