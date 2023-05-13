package main

import "fmt"

func main() {
	fmt.Println("0. Create array")
	my1stArray := [3]string{"Hola", "cara", "cola"}
	fmt.Println(my1stArray)

	fmt.Println("\n1. Classic loop")
	for i := 0; i < len(my1stArray); i++ {
		fmt.Printf("Position %d - Value %s \n", i, my1stArray[i])
	}

	fmt.Println("\n2. Range loop")
	// range create a copy of the array
	for i, value := range my1stArray {
		value := value + "-ggg"
		fmt.Printf("Position %d - Value %s \n", i, value)
	}
	fmt.Println(my1stArray)

	fmt.Println("\n3. Another way to create an array")
	my2ndArray := [...]string{"Hola", "cara", "cola"}
	fmt.Println("my2ndArray: ", my2ndArray)
	fmt.Printf("my2ndArray: %v \n", my2ndArray)
	for i, value := range my2ndArray {
		fmt.Printf("Position %d - Value %s \n", i, value)
	}

	fmt.Println("\n4. Arrays are copied")
	oneCopy := my2ndArray
	fmt.Println("oneCopy: ", oneCopy)
	oneCopy[0] = "Hello"
	fmt.Printf("oneCopy modified:  %v \n", oneCopy)
	fmt.Printf("my2ndArray: %v \n", my2ndArray)

	fmt.Println("\n5. Arrays passed by value")
	my3rdArray := [...]string{
		"Hola",
		"cara",
		"cola",
	}
	fmt.Println("my3rdArray before: ", my3rdArray)
	addText(my3rdArray)
	fmt.Println("my3rdArray after: ", my3rdArray)
}

func addText(vector [3]string) {
	for i := 0; i < len(vector); i++ {
		vector[i] = vector[i] + "-copy"
	}
	fmt.Println("Array copy: ", vector)
}
