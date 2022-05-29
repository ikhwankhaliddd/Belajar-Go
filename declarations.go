package main

import "fmt"

func main() {

	// type berfungsi untuk membuat alias dari sebuah tipe data
	
	type NoKTP string
	type Married bool

	var noKTPIkhwan NoKTP = "213861239"
	var marriedStatus Married = false
	fmt.Println(noKTPIkhwan)
	fmt.Println(marriedStatus)
}
