package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Ikhwan Khalid"

	fmt.Println(name)

	name = "Ikhwan Khalid"

	fmt.Println(name)

	var date = 25

	fmt.Println("Tanggal lahir = ", date)

	month := 3 // Deklarasi awal bisa menggunakan cara ini untuk menghindari kata kunci "var"

	fmt.Println("Bulan lahir = ", month)

	var (
		firstName = "Muhammad Ikhwan"
		lastName  = "Khaleed"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
