package gom

import "fmt"

type Element struct {
	Tag        string
	isFinite   bool
	attributes []Attribute
	children   []*Element
}

func H(t string, isFinite ...bool) *Element {
	_isFinite := false
	if len(isFinite) > 0 && isFinite[0] {
		_isFinite = true
	}

	return &Element{t, _isFinite, []Attribute{}, []*Element{}}
}

func (el *Element) Attrs(attrs ...Attribute) *Element {
	el.attributes = attrs
	return el
}

func (el *Element) Ch(children ...*Element) *Element {
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

	html += fmt.Sprintf("<%s%s>", el.Tag, attrs)

	for _, child := range el.children {
		html += child.Build()
	}

	html += fmt.Sprintf("</%s>", el.Tag)
	return html
}
