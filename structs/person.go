package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func newPerson(name string, age uint8) *Person {
	p := &Person{
		name: name,
		age:  age,
	}

	return p
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getAge() uint8 {
	return p.age
}

func main() {
	elChispas := &Person{
		name: "Juanito",
		age:  41,
	}
	fmt.Println(*elChispas)
	fmt.Println(elChispas.name)
	fmt.Println(elChispas.age)

	miki := newPerson("Miki", 43)
	fmt.Println(*miki)
	fmt.Println(miki.getName())
	fmt.Println(miki.getAge())
}
