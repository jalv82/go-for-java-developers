package main

import "fmt"

func main() {
	message := my1stFunction("Hello")
	fmt.Println(message)

	message, value := my2ndFunction("Hello")
	fmt.Printf("%s - %d \n", message, value)

	message, value = my3rdFunction("Hello", "world!", 4)
	fmt.Printf("%s - %d", message, value)
}

func my1stFunction(word string) string {
	return word + " world!"
}

func my2ndFunction(word string) (string, int) {
	return word + " world!", 8
}

func my3rdFunction(hello, world string, initValue int) (message string, value int) {
	message = fmt.Sprintf("%s %s", hello, world)
	value = initValue

	return message, value
}
