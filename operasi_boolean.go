package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusUjian bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusUjian && lulusAbsensi

	fmt.Println(lulus)
}