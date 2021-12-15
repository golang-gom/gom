package gom

import "fmt"

type Attribute struct {
	Name  string
	Value string
}

func Attr(name string, value string) Attribute {
	return Attribute{name, value}
}

func (a Attribute) Build() (attrs string) {
	attrs = fmt.Sprintf("%s=\"%s\"", a.Name, a.Value)
	return
}
