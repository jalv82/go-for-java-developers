package my_other_package

type PersonJava struct {
	name string `json:"nombre"`
	age  uint8  `json:"edad"`
}

func NewPersonJava(name string, age uint8) *PersonJava {
	p := &PersonJava{
		name: name,
		age:  age,
	}

	return p
}

func (p *PersonJava) GetName() string {
	return p.name
}

func (p *PersonJava) GetAge() uint8 {
	return p.age
}

type Person struct {
	Name string `json:"nombre"`
	Age  uint8  `json:"edad"`
}

func NewPerson(name string, age uint8) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}

	return p
}
