package main

import (
	"fmt"
)

func main() {
	fmt.Println("0. Create map")
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Printf("Content: %v. Type: %T \n", myMap, myMap)

	fmt.Println("\n1. Range loop")
	for k, v := range myMap {
		fmt.Printf("key: %s - Value: %d\n", k, v)
	}

	fmt.Println("\n2. Add item to map")
	myMap["four"] = 4
	fmt.Printf("Content: %v. Type: %T \n", myMap, myMap)

	fmt.Println("\n3. Delete item")
	delete(myMap, "two")
	fmt.Printf("Content: %v. Type: %T \n", myMap, myMap)
	if value, ok := myMap["two"]; ok {
		fmt.Println("Value is: ", value)
	} else {
		fmt.Println("Key not found")
	}

	fmt.Println("\n4. Maps aren't copied")
	fmt.Printf("myMap content: %v. Type: %T \n", myMap, myMap)
	myCopyMap := myMap
	fmt.Printf("myCopyMap content: %v. Type: %T \n", myCopyMap, myCopyMap)
	delete(myMap, "four")
	fmt.Printf("myMap content: %v. Type: %T \n", myMap, myMap)
	fmt.Printf("myCopyMap content: %v. Type: %T \n", myCopyMap, myCopyMap)
}
