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

		result := H("h1").A(Attr("class", "title"))
		expected := Element{Tag: "h1", attributes: []Attribute{{Name: "class", Value: "title"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("With Children", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").C(H("span"))
		expected := Element{Tag: "h1", children: []*Element{{Tag: "span"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("With Attrs and Children", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").A(Attr("class", "title")).C(H("span"))
		expected := Element{Tag: "h1", attributes: []Attribute{{Name: "class", Value: "title"}}, children: []*Element{{Tag: "span"}}}

		must.Equal(result.Tag, expected.Tag)
		must.Equal(len(result.attributes), len(expected.attributes))
		must.Equal(len(result.children), len(expected.children))
	})

	t.Run("Build", func(t *testing.T) {
		must := must.New(t)

		result := H("h1").A(Attr("class", "title")).C(H("welcome", IsFinite))
		expected := "<h1 class=\"title\">welcome</h1>"

		must.Equal(result.Build(), expected)
	})

	t.Run("Build With No Close", func(t *testing.T) {
		must := must.New(t)

		result := H("hr", NoClose)
		expected := "<hr />"

		must.Equal(result.Build(), expected)
	})

}
