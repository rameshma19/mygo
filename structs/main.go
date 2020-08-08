package main

import "fmt"

type person struct {
	fn      string
	ln      string
	age     int
	gender  bool
	contact struct {
		flatNo  string
		aptName string
		area    string
		city    string
		pincode int
	}
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateFn(newfn string) {
	(*pointerToPerson).fn = newfn
}

func main() {
	var p1 person

	fmt.Println(p1)

	p1.fn = "Ramesh"
	p1.ln = "Mantripragada"
	p1.age = 45
	p1.gender = true
	p1.contact.flatNo = "5025"
	p1.contact.aptName = "Shobha"
	p1.contact.area = "Harlur Rd"
	p1.contact.city = "bangalore"
	p1.contact.pincode = 5025

	//p1Pointer := &p1
	p1.updateFn("Aarav")
	p1.print()
}
