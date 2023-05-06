package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration: ", i)
	}

	// ++i not supported
	for i := 0; i < 10; i++ {
		fmt.Printf("Other iteration: %d \n", i)
	}
}
