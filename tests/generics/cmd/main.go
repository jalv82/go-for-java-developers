package main

import (
	"generics"
	"log"
)

func main() {
	log.Println("# Sum with Generics")
	log.Println("Integers: ", generics.Sum(3, 4, 5, 6))
	log.Println("Floats: ", generics.Sum(3.6, 4.5, 5.4, 6.3))
}
