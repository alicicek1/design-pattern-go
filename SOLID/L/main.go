package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, " but got ", actualArea, "\n")
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.height = size
	sq.width = size
	return &sq
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
