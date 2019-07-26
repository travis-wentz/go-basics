package main

import "fmt"

func main(){
	//slices are basically like arraylists in Java, but they're called slices because that's so much cooler
	//instead of being reasonable and having 'slice' be the key word to make a slice, they decided to make the key word be 'make'
	//you can instantiate with a non-zero size
	elves := make([]string, 2)
	elves[0] = "Elrond"
	elves[1] = "Arwen"
	
	//and you can append to the arrayli... to the slice
	elves = append(elves, "Celeborn")
	fmt.Println(elves)

	//you can also instantiate with values
	hobbits := []string{"Bilbo", "Frodo", "Samwise the Brave", "Merry", "Pippin"}
	fmt.Println(hobbits)

	//and they're easy to copy
	elvesCopy := make([]string, len(elves))
	copy(elvesCopy, elves)
	fmt.Println(elvesCopy)
}
