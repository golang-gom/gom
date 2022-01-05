package gom

func HTML() *Element { return H("html") }
func Head() *Element { return H("head") }
func Body() *Element { return H("body") }
func Div() *Element  { return H("div") }

func Table() *Element { return H("table") }
func THead() *Element { return H("thead") }
func TBody() *Element { return H("tbody") }
func TR() *Element    { return H("tr") }
func TD() *Element    { return H("td") }

func Title() *Element { return H("title") }
func H1() *Element    { return H("h1") }
func H2() *Element    { return H("h2") }
func H3() *Element    { return H("h3") }
func H4() *Element    { return H("h4") }
func H5() *Element    { return H("h5") }
func Span() *Element  { return H("span") }
func I() *Element     { return H("i") }
func B() *Element     { return H("b") }
func U() *Element     { return H("u") }

func Text(text string) *Element { return H(text, IsFinite) }
