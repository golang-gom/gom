package gom

import (
	"fmt"
	"os"
)

type Element struct {
	Tag        string
	isFinite   bool
	noClose    bool
	attributes []Attribute
	children   []*Element
}

func H(t string, options ...*Option) *Element {
	el := &Element{Tag: t, attributes: []Attribute{}, children: []*Element{}}
	for _, option := range options {
		if option.Name == IsFinite.Name {
			el.isFinite = true
		} else if option.Name == NoClose.Name {
			el.noClose = true
		}
	}

	return el
}

func (el *Element) A(attrs ...Attribute) *Element {
	el.attributes = attrs
	return el
}

func (el *Element) C(children ...*Element) *Element {
	el.children = children
	return el
}

func (el Element) Build() (html string) {
	if el.isFinite {
		html = el.Tag
		return
	}

	attrs := ""
	for _, attr := range el.attributes {
		attrs += " " + attr.Build()
	}

	noClose := ""
	if el.noClose {
		noClose = " /"
	}
	html += fmt.Sprintf("<%s%s%s>", el.Tag, attrs, noClose)

	for _, child := range el.children {
		html += child.Build()
	}

	if !el.noClose {
		html += fmt.Sprintf("</%s>", el.Tag)
	}
	return html
}

func (el Element) Export(loc string) (err error) {
	f, err := os.Create(loc)
	if err != nil {
		return
	}
	defer f.Close()

	htmlString := el.Build()
	_, err = f.WriteString(htmlString)
	return
}
