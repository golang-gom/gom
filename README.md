# gom
DOM building using Go

# Usage
```go
package main

import "github.com/hadihammurabi/gom"

func main() {
	dom := gom.H("html").Children(
		gom.H("head").Children(
			gom.H("title").Children(
				gom.H("Home Page", true),
			),
		),
		gom.H("body").Children(
			gom.H("h1").Children(
				gom.H("Welcome to our Home Page!", true),
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
        dom := gom.HTML.Children(
                gom.Head.Children(
                        gom.Title.Children(
                                gom.Text("Home Page"),
                        ),
                ),
                gom.Body.Children(
                        gom.H1.Children(
                                gom.Text("Welcome to our Home Page!"),
                        ),
                ),
        )

        println(dom.Build())
}
```
