package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"Monkey", "D", "Luffy"}
	printMessage("Shashiburinara", names)
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
	randomValue = randomWithRange(10, 100)
	fmt.Println("Random Number:", randomValue)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	values := rand.Int()%(max-min+1) - min
	return values
}
