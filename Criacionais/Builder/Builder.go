package main

import "fmt"

// Produto que queremos construir
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Builder para Person
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (b *PersonBuilder) FirstName(firstName string) *PersonBuilder {
	b.person.FirstName = firstName
	return b
}

func (b *PersonBuilder) LastName(lastName string) *PersonBuilder {
	b.person.LastName = lastName
	return b
}

func (b *PersonBuilder) Age(age int) *PersonBuilder {
	b.person.Age = age
	return b
}

func (b *PersonBuilder) Email(email string) *PersonBuilder {
	b.person.Email = email
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	person := NewPersonBuilder().
		FirstName("Marcelo").
		LastName("Cerqueira").
		Age(35).
		Email("marcelo.raposo@gmail.com").
		Build()

	fmt.Printf("%+v\n", person)
}
