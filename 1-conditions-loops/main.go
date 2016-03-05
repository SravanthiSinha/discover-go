package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomNumber = rand.Intn(100)
	if randomNumber > 50 {
		fmt.Println("my random number is", randomNumber, "and is greater than 50")
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	} else {

	}
	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	} else {

	}
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for _, value := range holbertonFounders {
		fmt.Println(value, "is a founder")
	}
}
