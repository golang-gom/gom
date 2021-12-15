# gom
DOM building using Go

# Usage
```go
package main

import "github.com/hadihammurabi/gom"

func main() {
	dom := gom.H("html").C(
		gom.H("head").C(
			gom.H("title").C(
				gom.H("Home Page", gom.IsFinite),
			),
		),
		gom.H("body").C(
			gom.H("h1").C(
				gom.H("Welcome to our Home Page!", gom.IsFinite),
			),
		),
	)

	println(dom.Build())
}
```

it will shows
```html
<html><head><title>Home Page</title></head><body><h1>Welcome to our Home Page!</h1></body></html>
```

or use DOM utilities
```go
package main

import "github.com/hadihammurabi/gom"

func main() {
	dom := gom.HTML.C(
		gom.Head.C(
			gom.Title.C(
				gom.Text("Home Page"),
			),
		),
		gom.Body.C(
			gom.H1.C(
				gom.Text("Welcome to Home Page!"),
			),
		),
	)

	println(dom.Build())
}
```
