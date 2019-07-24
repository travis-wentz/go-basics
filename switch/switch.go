package main

import "fmt"

func main() {
	//breaks not required in go switch statements because go is not stupid
	team := "brown"
	switch team {
	case "purple":
		fmt.Println("Did you hear what that purple-belly just called the ship?")
	case "brown":
		fmt.Println("Old browncoats, eh?")
	case "blue":
		fmt.Println("Two by two, hands of blue.")
	}
}
