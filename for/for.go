package main

import "fmt"

func main() {
	
	//like a while loop
	preFib := 0
	fib := 1
	fmt.Printf("%d ", preFib)
	for fib <= 144 {
		fmt.Printf("%d ", fib)
		fib += preFib
		preFib = fib - preFib
	}
	fmt.Println()

	//like a traditional for loop
	preFib = 0
	fib = 1
	fmt.Printf("%d ", preFib)
	for i := 0; i < 12; i++ {
		fmt.Printf("%d ", fib)
		fib += preFib
		preFib = fib - preFib
	}
	fmt.Println()

	//jumping to the next iteration of the for loop
	preFib = 0
	fib = 1
	for i := 0; i <= 12; i++ {
		if i == 0 {
			fmt.Print("0 ")
			continue
		}
		fmt.Printf("%d ", fib)
		fib += preFib
		preFib = fib - preFib
	}
	fmt.Println()

	//for loop without condition
	preFib = 0
	fib = 1
	fmt.Printf("%d ", preFib)
	for {
		fmt.Printf("%d ", fib)
		fib += preFib
		preFib = fib - preFib	
		if fib < 145 {
			continue
		}
		break
	}
	fmt.Println()
}

