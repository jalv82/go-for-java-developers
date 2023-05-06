package main

import "fmt"

func main() {
	condition := "java"
	switch condition {
	case "java", "jee":
		fmt.Printf("You choose %s language\n", condition)
	case "go":
		fmt.Printf("You choose %s language\n", condition)
	default:
		fmt.Printf("You are rust developer\n")
	}

	condition = "gof"
	switch {
	case condition == "java", condition == "jee":
		fmt.Printf("You choose %s language", condition)
	case condition == "go":
		fmt.Printf("You choose %s language", condition)
	default:
		fmt.Printf("You are rust developer 🦀")
	}
}
