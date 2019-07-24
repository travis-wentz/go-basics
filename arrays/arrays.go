package main

import "fmt"

func main() {
	var crew = [9]string{"Mal", "Zoe", "Wash", "Jayne", "Book", "Inara", "Kaylee", "Simon", "River"}
	fmt.Println(crew)

	var associates [4]string
	associates[0] = "Badger"
	associates[1] = "Mr. Universe"
	associates[2] = "Fanty"
	associates[3] = "Mingo"
	fmt.Println(associates)
}
