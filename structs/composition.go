package main

import (
	"encoding/json"
	"fmt"
	"structs/my_other_package"

	address "structs/my_package"
)

type PersonComposed struct {
	Name string `json:"nombre"`
	Age  uint8  `json:"edad"`
	// Street string
	address.Address
}

func main() {
	fmt.Println("0. Create struct with Java style")
	elChispasJava := my_other_package.NewPersonJava("Juanito", 41)
	fmt.Println("struct: ", *elChispasJava)
	fmt.Println("name: ", elChispasJava.GetName())
	fmt.Println("age: ", elChispasJava.GetAge())

	fmt.Println("\n1. Convert struct to json")
	jsonBytes, err := json.Marshal(elChispasJava)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Struct: ", elChispasJava)
	fmt.Println("JSON: ", string(jsonBytes))

	fmt.Println("\n2. Convert json to struct")
	myJson := `{"nombre":"Albert","edad":21, "street":"Into the station", "number": 13}`
	fmt.Println("JSON: ", myJson)
	var carl my_other_package.PersonJava
	err = json.Unmarshal([]byte(myJson), &carl)
	if err != nil {
		return
	}
	fmt.Println("Struct: ", carl)

	fmt.Println("\n3. Create struct like Go style")
	elChispas := &PersonComposed{
		Name: "Juanito",
		Age:  41,
		// Street: "street persona no de address",
		Address: address.Address{
			Street: "Under the bridge",
			Number: 29,
		},
	}
	fmt.Println("struct: ", *elChispas)
	fmt.Println("name: ", elChispas.Name)
	fmt.Println("age: ", elChispas.Age)
	fmt.Println("street: ", elChispas.Street)
	fmt.Println("number: ", elChispas.Address.Number)

	fmt.Println("\n4. Create other struct by constructor")
	miki := my_other_package.NewPerson("Miki", 43)
	fmt.Println("struct: ", *miki)
	fmt.Println("name: ", miki.Name)
	fmt.Println("age: ", miki.Age)

	fmt.Println("\n5. Convert struct to json")
	jsonBytes, err = json.Marshal(elChispas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Struct: ", elChispas)
	fmt.Println("JSON: ", string(jsonBytes))

	fmt.Println("\n6. Convert json to struct")
	myJson = `{"nombre":"Albert","edad":21, "street":"Into the station", "number": 13}`
	fmt.Println("JSON: ", myJson)
	albert := PersonComposed{}
	err = json.Unmarshal([]byte(myJson), &albert)
	if err != nil {
		return
	}
	fmt.Println("Struct: ", albert)
	fmt.Println("Struct: ", albert.Street)
	fmt.Println("Struct: ", albert.Address.Street)
}
