package tagstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TagStruct struct {
	ID          int   `ts:"name=id"`
	CategoryID  int32 `ts:"name=category_id"`
	Name        string
	Description string   `ts:"name=description"`
	Keyword     string   `ts:"name=keyword"`
	Other       []string `ts:"name=other"`
	Required    bool     `ts:"name=required"`
	Float32     float32  `ts:"name=float32"`
	Float64     float64  `ts:"name=float64"`
}

func TestTagStruct(t *testing.T) {
	t.Run("test struct definition", func(t *testing.T) {
		assert.NotPanics(t, func() {
			defined := New(TagStruct{})
			assert.NotNil(t, defined)
			result, err := defined.ParseTag("description='hello \\'world',keyword='something else',other=['a','b'],id=3,category_id=42,required=true,float32=1.1,float64=2.2")
			assert.Nil(t, err)
			assert.Equal(t, "hello 'world", result.Description)
			assert.Equal(t, "something else", result.Keyword)
			assert.Equal(t, 3, result.ID)
			assert.Equal(t, int32(42), result.CategoryID)
			assert.Equal(t, []string{"a", "b"}, result.Other)
			assert.True(t, result.Required)
		})
	})

	t.Run("test parse struct", func(t *testing.T) {
		type TestStruct struct {
			ID int `custom:"id=42,required"`
		}

		m, err := New(TagStruct{}).ParseStruct(TestStruct{}, "custom")
		assert.Nil(t, err)
		assert.NotZero(t, m)
	})
}
