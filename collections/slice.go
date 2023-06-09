package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("0. Create array")
	myArray := [...]string{"Hola", "Cara", "Cola"}
	fmt.Printf("Content: %v. Type: %T \n", myArray, myArray)

	fmt.Println("\n1. Convert array to slice")
	mySlice := myArray[:]
	fmt.Printf("Content: %v. Type: %T \n", mySlice, mySlice)
	fmt.Println("Length: ", len(mySlice), "Capacity: ", cap(mySlice))

	fmt.Println("\n2. Append one item to the slice, change its capacity")
	mySlice = append(mySlice, "!")
	fmt.Println("Slice: ", mySlice, "Length: ", len(mySlice), "Capacity: ", cap(mySlice))

	fmt.Println("\n3. Print positions")
	fmt.Println("Slice: ", mySlice[:2])
	fmt.Println("Slice: ", mySlice[2:3])
	for counter, value := range mySlice {
		fmt.Printf("Position %d - Value %s \n", counter, value)
	}

	fmt.Println("\n4. Order asc")
	sort.StringSlice(mySlice).Sort()
	for counter, value := range mySlice {
		fmt.Printf("Position %d - Value %s \n", counter, value)
	}

	fmt.Println("\n5. Order desc")
	sort.Sort(sort.Reverse(sort.StringSlice(mySlice)))
	for counter, value := range mySlice {
		fmt.Printf("Position %d - Value %s \n", counter, value)
	}

	fmt.Println("\n6. Order desc (slice of int)")
	my1stSlice := make([]uint, 0, 6)
	my1stSlice = append(my1stSlice, 11, 13, 17, 19, 23, 29)

	prefix2 := my1stSlice[4:6]
	fmt.Println("Original prefix2: ", prefix2)

	my1stSlice[4] = 50
	fmt.Println("Change in my1stSlice: ", prefix2)

	prefix2[0] = 100
	fmt.Println("Change in prefix2: ", my1stSlice)

}
