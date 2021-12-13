package gom

import (
	"testing"

	"github.com/golang-must/must"
)

func TestAttr(t *testing.T) {

	t.Run("Attr", func(t *testing.T) {
		must := must.New(t)

		result := Attr("class", "title")
		expected := Attribute{"class", "title"}

		must.Equal(result.Name, expected.Name)
		must.Equal(result.Value, expected.Value)
	})

}
