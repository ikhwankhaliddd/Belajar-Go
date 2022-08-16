package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Ikhwan",
		"adress": "Siantar",
	}
	person["title"] = "Software Engineer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ikhwan"
	book["Ups"] = "Salah"
	fmt.Println(book)

	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
