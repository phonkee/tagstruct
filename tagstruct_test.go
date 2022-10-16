package tagstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExampleStruct struct {
	ID          int   `ts:"name=id"`
	CategoryID  int32 `ts:"name=category_id"`
	Name        string
	Description string   `ts:"name=description"`
	Keyword     string   `ts:"name=keyword"`
	Other       []string `ts:"name=other"`
}

func TestTagStruct(t *testing.T) {
	t.Run("test struct definition", func(t *testing.T) {
		assert.NotPanics(t, func() {
			defined := New(ExampleStruct{})
			assert.NotNil(t, defined)
			result, err := defined.ParseTag("description='hello world',keyword='something',other=['a','b'],id=3,category_id=42")
			assert.Nil(t, err)
			assert.Equal(t, "hello world", result.Description)
			assert.Equal(t, "something", result.Keyword)
			assert.Equal(t, 3, result.ID)
			assert.Equal(t, int32(42), result.CategoryID)
			assert.Equal(t, []string{"a", "b"}, result.Other)
		})
	})
}
