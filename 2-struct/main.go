package main

import (
	"fmt"
	"time"
)

type user struct {
    Name string `json:"name"`
    DOB string `json:"date_of_birth"`
    City string `json:"city"`
}

func main() {
	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
   	fmt.Printf("Hello %s\n", u.Name)
   	layout := "January 2, 2006"
   	y, _ := time.Parse(layout, u.DOB)
   	t := time.Now()
   	age := (t.Year()-y.Year())
	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, age)
}
