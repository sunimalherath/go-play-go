package oop

import "fmt"

type Person struct {
	Name string
}

// Person function receiver
func (p Person) PrintInfo() {
	fmt.Printf("Person Name: %s\n", p.Name)
}

// Person pointer receiver
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// Person receiver - without pointer
func (p Person) ChangeNameNoPointer(newName string) {
	p.Name = newName
}

