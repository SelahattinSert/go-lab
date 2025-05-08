package main

import "fmt"

type contactInfo struct {
	email string
	zip   string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jane := person{
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email: "jane@gmail.com",
			zip:   "12345",
		},
	}

	jane.setName("Janet")
	fmt.Println(jane)
}

func (pointerToPerson *person) setName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Println("%+v", p)
}
