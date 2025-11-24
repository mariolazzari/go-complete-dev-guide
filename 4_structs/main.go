package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mario := person{"Mario", "Lazzari"}
	fmt.Println(mario)
	maria := person{firstName: "Maria", lastName: "Lazzari"}
	fmt.Println(maria)

}
