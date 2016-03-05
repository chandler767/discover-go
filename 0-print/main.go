package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello, we are Holberton School\n")
	t := time.Now()
	fmt.Printf("the date is %s\n", t)
	fmt.Printf("the year is %d\n", t.Year())
	fmt.Printf("the month is %s\n", t.Month())
	fmt.Printf("the day is %d\n", t.Day())
}