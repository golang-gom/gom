package gom

var (
	HTML = func() *Element { return H("html") }
	Head = func() *Element { return H("head") }
	Body = func() *Element { return H("body") }
	Div  = func() *Element { return H("div") }

	Table = func() *Element { return H("table") }
	THead = func() *Element { return H("thead") }
	TBody = func() *Element { return H("tbody") }
	TR    = func() *Element { return H("tr") }
	TD    = func() *Element { return H("td") }

	Title = func() *Element { return H("title") }
	H1    = func() *Element { return H("h1") }
	H2    = func() *Element { return H("h2") }
	H3    = func() *Element { return H("h3") }
	H4    = func() *Element { return H("h4") }
	H5    = func() *Element { return H("h5") }
	Span  = func() *Element { return H("span") }
	I     = func() *Element { return H("i") }
	B     = func() *Element { return H("b") }
	U     = func() *Element { return H("u") }
)

func Text(text string) *Element {
	return H(text, IsFinite)
}
