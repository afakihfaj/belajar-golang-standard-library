package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	Hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(Hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
