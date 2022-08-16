package main

import "fmt"

func main() {

	name := "Nasution"

	if name == "Ikhwan" {
		fmt.Println("Hello " + name)
	} else if name == "Nasution" {
		fmt.Println("Hello Bang " + name)
	} else {
		fmt.Println("Hello " + name)
	}

	// Short statement if

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
