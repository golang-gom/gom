package gom

var (
	HTML = H("html")
	Head = H("head")
	Body = H("body")
	Div  = H("div")

	Table = H("table")
	THead = H("thead")
	TBody = H("tbody")
	TR    = H("tr")
	TD    = H("td")

	Title = H("title")
	H1    = H("h1")
	H2    = H("h2")
	H3    = H("h3")
	H4    = H("h4")
	H5    = H("h5")
	Span  = H("span")
	I     = H("i")
	B     = H("b")
	U     = H("u")
)

func Text(text string) *Element {
	return H(text, IsFinite)
}
