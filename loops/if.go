package main

import "fmt"

func main() {
	condition := "go"

	if condition == "java" || condition == "jee" {
		fmt.Printf("You choose %s language", condition)
	} else if condition == "go" {
		fmt.Printf("You choose %s language", condition)
	} else {
		fmt.Printf("You are rust developer 🦀")
	}
}
