package main

import "fmt"

type Person struct {
	//address
	StreetAddress, PostCode, City string

	//Job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*p}
}

func (p *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*p}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (p *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	p.person.CompanyName = companyName
	return p
}

func (p *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	p.person.Position = position
	return p
}

func (p *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	p.person.AnnualIncome = annualIncome
	return p
}

func (p *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	p.person.StreetAddress = streetAddress
	return p
}

func (p *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	p.person.City = city
	return p
}

func (p *PersonAddressBuilder) WithPostCode(postCode string) *PersonAddressBuilder {
	p.person.PostCode = postCode
	return p
}

type PersonAddressBuilderInterface interface {
	At(streetAddress string) *PersonAddressBuilder
	In(city string) *PersonAddressBuilder
	WithPostCode(postCode string) *PersonAddressBuilder
}

type PersonJobBuilderInterface interface {
	At(companyName string) *PersonJobBuilder
	AsA(position string) *PersonJobBuilder
	Earning(annualIncome int) *PersonJobBuilder
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London").
		In("London").
		WithPostCode("123qwe").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)

	person := pb.Build()
	fmt.Println(person)
}
