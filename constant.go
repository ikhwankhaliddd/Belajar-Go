package main

func main() {
	const firstName string = "Ikhwan"
	const lastName = "Khaleed"
	const value = 1000

	// cara lain membuat constant adalah dengan multiple constant

	const (
		firstNamee string = "Ikhwan"
		lastNamee         = "Khaleed"
		valuee            = 1000
	)

	//ini bakal error karena const

	firstName = "Tidak bisa diubah karena constant"
	lastName = "Tidak bisa diubah karena constant"

}
