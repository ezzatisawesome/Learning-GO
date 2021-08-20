package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("You can ride!")
	} else {
		fmt.Println("You cannot ride")
		fmt.Printf("Wait %d years", 18-age)
	}
}
