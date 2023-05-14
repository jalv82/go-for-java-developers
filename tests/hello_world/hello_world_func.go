package main

import "fmt"

func main() {
	message := my1stFunction("Hello")
	fmt.Println(message)
}

func my1stFunction(word string) string {
	return word + " world!"
}
