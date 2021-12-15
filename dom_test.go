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
