package main

import "fmt"

func main() {
	shp := map[string]string{
		"Captain":      "Luffy",
		"Vice Captain": "Zorro",
		"Doctor":       "Chopper",
	}
	//fmt.Println(shp)
	//fmt.Println(shp["Doctor"])

	for key, val := range shp {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println(len(shp))
	delete(shp, "Doctor")
	fmt.Println(len(shp))
	for key, val := range shp {
		fmt.Println(key, " \t:", val)
	}
	var value, isExist = shp["Navigator"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Data is not Exist")
	}
	var strawHatPirate = []map[string]string{
		{"Nama": "Luffy", "Jabatan": "Captain"},
		{"Nama": "Zorro", "Jabatan": "Vice Captain"},
		{"Nama": "Sanji", "Jabatan": "Koki"},
	}

	for _, shp := range strawHatPirate {
		fmt.Println(shp["Nama"], " \t:", shp["Jabatan"])
	}
}
