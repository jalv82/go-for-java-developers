package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 1. Read csv file
	file, err := os.Open("cars.csv")
	if err != nil {
		log.Panicln("Can't read the file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Panicln("Can't close the file")
		}
	}()

	reader := csv.NewReader(file)
	lines, _ := reader.ReadAll()

	// 2. Process each line
	brands := make(map[string]int)
	for _, line := range lines {
		// 3. Getting data of each line
		// Ford 9000 23
		items := strings.Split(line[0], " ")
		quantity, _ := strconv.Atoi(items[1])
		brands[items[0]] = quantity
	}

	quantities := make([]string, 0, len(brands))
	total := 0
	for key, value := range brands {
		quantities = append(quantities, key)
		total += value
	}

	// 4. Order by quantity asc -> desc
	log.Println("Order by quantity asc -> desc: ")
	sort.SliceStable(quantities, func(i, j int) bool {
		return brands[quantities[i]] > brands[quantities[j]]
	})
	for _, brand := range quantities {
		quantity := brands[brand]
		fmt.Println(brand, quantity)
	}

	// 5. Order by quantity desc -> asc
	log.Println("Order by quantity desc -> asc: ")
	sort.SliceStable(quantities, func(i, j int) bool {
		return brands[quantities[i]] < brands[quantities[j]]
	})
	for _, brand := range quantities {
		quantity := brands[brand]
		fmt.Println(brand, quantity)
	}

	// 6. Add total quantity
	log.Println("Total quantity: ", total)
}
