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

func main() {
	kelton := person{
		firstName: "Kelton",
		lastName:  "Williams",
		contact: contactInfo{
			email:   "kelton@gmail.com",
			zipCode: 1,
		},
	}

	kelton.updateName("Kelton II")
	kelton.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
