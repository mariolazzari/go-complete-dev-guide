package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%s after assignment: %+v\n", p.firstName, p)
}

func main() {
	mario := person{"Mario", "Lazzari", contactInfo{"mario.lazzari@gmail.com", 38066}}
	mario.print()
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	maria.print()

	var alex person
	fmt.Printf("alxe defaults: %+v\n", alex)
	alex.firstName = "Alex"
	alex.lastName = "Lifeson"
	alex.print()

	mary := person{
		firstName: "Mary",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "mary@mail.com",
			zipCode: 12345,
		},
	}
	mary.print()

	mary.updateName("Marianne")
	mary.print()
}
