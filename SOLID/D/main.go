package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

//low-level module
type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, relation := range r.relations {
		if relation.relationship == Parent && relation.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

//High-Level Module
type Research struct {
	//break DIP
	//relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	//relations := r.relationships.relations
	//for _, relation := range relations {
	//	if relation.from.name == "John" && relation.relationship == Parent {
	//		fmt.Println("John has a child called ", relation.to.name)
	//	}
	//}

	for _, person := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", person.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	re := Relationships{}
	re.AddParentAndChild(&parent, &child1)
	re.AddParentAndChild(&parent, &child2)

	r := Research{&re}
	r.Investigate()
}
