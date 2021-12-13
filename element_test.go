package gom

import (
	"testing"

	"github.com/golang-must/must"
)

func TestH(t *testing.T) {

	t.Run("H", func(t *testing.T) {
		must := must.New(t)
		result := H("h1")
		expected := Element{Tag: "h1"}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("With Attrs", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").Attrs(Attr("class", "title"))
		expected := Element{Tag: "h1", attributes: []Attribute{{Name: "class", Value: "title"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("With Children", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").Children(H("span"))
		expected := Element{Tag: "h1", children: []*Element{{Tag: "span"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("With Attrs and Children", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").Attrs(Attr("class", "title")).Children(H("span"))
		expected := Element{Tag: "h1", attributes: []Attribute{{Name: "class", Value: "title"}}, children: []*Element{{Tag: "span"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("Build", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").Attrs(Attr("class", "title")).Children(H("welcome", true))
		expected := "<h1>welcome</h1>"

		must.Equal(result.Build(), expected)
	})

}
