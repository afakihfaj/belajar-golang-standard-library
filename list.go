package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Afakih")
	data.PushBack("Fajduwani")
	data.PushBack("Dewangga")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // afakih

	next := head.Next() // fajduwani
	fmt.Println(next.Value)

	next = next.Next() // dewangga
	fmt.Println(next.Value)
}
