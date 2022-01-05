package gom

import (
	"testing"

	"github.com/golang-must/must"
)

func TestText(t *testing.T) {

	t.Run("Text", func(t *testing.T) {
		must := must.New(t)

		result := Text("test")
		expected := H("test", IsFinite)

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.isFinite, expected.isFinite)
	})

}

func TestDOMUtils(t *testing.T) {

	t.Run("HTML", func(t *testing.T) {
		must := must.New(t)

		result := HTML()
		expected := H("html")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Head", func(t *testing.T) {
		must := must.New(t)

		result := Head()
		expected := H("head")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Body", func(t *testing.T) {
		must := must.New(t)

		result := Body()
		expected := H("body")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Div", func(t *testing.T) {
		must := must.New(t)

		result := Div()
		expected := H("div")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Table", func(t *testing.T) {
		must := must.New(t)

		result := Table()
		expected := H("table")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("THead", func(t *testing.T) {
		must := must.New(t)

		result := THead()
		expected := H("thead")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("TBody", func(t *testing.T) {
		must := must.New(t)

		result := TBody()
		expected := H("tbody")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("TR", func(t *testing.T) {
		must := must.New(t)

		result := TR()
		expected := H("tr")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("TD", func(t *testing.T) {
		must := must.New(t)

		result := TD()
		expected := H("td")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Title", func(t *testing.T) {
		must := must.New(t)

		result := Title()
		expected := H("title")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("H1", func(t *testing.T) {
		must := must.New(t)

		result := H1()
		expected := H("h1")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("H2", func(t *testing.T) {
		must := must.New(t)

		result := H2()
		expected := H("h2")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("H3", func(t *testing.T) {
		must := must.New(t)

		result := H3()
		expected := H("h3")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("H4", func(t *testing.T) {
		must := must.New(t)

		result := H4()
		expected := H("h4")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("H5", func(t *testing.T) {
		must := must.New(t)

		result := H5()
		expected := H("h5")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("Span", func(t *testing.T) {
		must := must.New(t)

		result := Span()
		expected := H("span")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})
	t.Run("I", func(t *testing.T) {
		must := must.New(t)

		result := I()
		expected := H("i")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("B", func(t *testing.T) {
		must := must.New(t)

		result := B()
		expected := H("b")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

	t.Run("U", func(t *testing.T) {
		must := must.New(t)

		result := U()
		expected := H("u")

		must.Equal(result.Tag, expected.Tag)
		must.Equal(result.Build(), expected.Build())
	})

}
