package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (h *HtmlElement) String() string {
	return h.string(0)
}

func (h *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, h.name))
	if len(h.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(h.text)
		sb.WriteString("\n")
	}

	for _, element := range h.elements {
		sb.WriteString(element.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, h.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{
		name:     rootName,
		text:     "",
		elements: []HtmlElement{},
	}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func main() {
	hello := "Hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"Hello", "World"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, word := range words {
		sb.WriteString("<li>")
		sb.WriteString(word)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	//
	fmt.Println()
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "Hello")
	b.AddChild("li", "World")
	fmt.Println(b.String())
}
