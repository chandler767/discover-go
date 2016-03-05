package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	rand.Seed(time.Now().Unix())
	randomNumber := random(1, 100)
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Printf("I am a student of the Holberton School\n")
	} else {
		fmt.Printf("I am not a student of the Holberton School\n")
	}
	beautifulWeather := true
	if beautifulWeather == true {
		fmt.Printf("It's a beautiful weather!\n")
	} else {
		fmt.Printf("It's not beautiful weather!\n")
	}
	holbertonFounders := make([]string, 3)
	holbertonFounders[0] = "Rudy Rigot"
	holbertonFounders[1] = "Sylvain Kalache,"
	holbertonFounders[2] = "Julien Barbier"
	for i := 0; i <= 2; i++ {
        fmt.Printf("%s is a founder\n", holbertonFounders[i])
    }
}