package main

import "fmt"

//OCP
//open for extension, closed for modification
//Enterprise pattern - Specification

type Color int
type Size int

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size && product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

/////////// Better solution

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}
type SizeSpecification struct {
	size Size
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type BetterFilter struct{}

func (b *BetterFilter) Filter(products []Product, specification Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if specification.IsSatisfied(&product) {
			result = append(result, &products[i])
		}
	}
	return result
}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"apple", green, small}
	tree := Product{"tree", green, large}
	house := Product{"house", blue, large}

	productList := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, product := range f.FilterByColor(productList, green) {
		fmt.Printf(" - %s is green\n", product.name)
	}

	/////////// Better solution
	fmt.Printf("Green products (new):\n")
	colorSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, product := range bf.Filter(productList, &colorSpec) {
		fmt.Printf(" - %s is green\n", product.name)
	}

	fmt.Printf("Green and large products (new):\n")
	sizeSpec := SizeSpecification{large}
	andSpec := AndSpecification{&colorSpec, &sizeSpec}
	for _, product := range bf.Filter(productList, &andSpec) {
		fmt.Printf(" - %s is green and large\n", product.name)
	}
}
