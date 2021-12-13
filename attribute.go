package gom

type Attribute struct {
	Name  string
	Value string
}

func Attr(name string, value string) Attribute {
	return Attribute{name, value}
}
