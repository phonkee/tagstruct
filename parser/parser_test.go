package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](t T) *T {
	return &t
}

func TestParse(t *testing.T) {
	t.Run("test parse simple values", func(t *testing.T) {
		data := []struct {
			input  string
			expect []Property
		}{
			{"", []Property{}},
			{"hello=1", []Property{
				{Position: 0, Name: "hello", Value: &Value{Position: 6, Number: ptr("1")}},
			}},
			{"hello='world'", []Property{
				{Position: 0, Name: "hello", Value: &Value{Position: 6, String: ptr("world")}},
			}},
			{"hello=world, hello= hello,  world=true, obj(one=1, two='two')", []Property{
				{Position: 0, Name: "hello", Value: &Value{Position: 6, String: ptr("world")}},
				{Position: 13, Name: "hello", Value: &Value{Position: 20, String: ptr("hello")}},
				{Position: 28, Name: "world", Value: &Value{Position: 34, String: ptr("true")}},
				{Position: 40, Name: "obj", Object: []Property{
					{Position: 44, Name: "one", Value: &Value{Position: 48, Number: ptr("1")}},
					{Position: 51, Name: "two", Value: &Value{Position: 55, String: ptr("two")}},
				}},
			}},
		}

		for _, item := range data {
			result, err := Parse(strings.NewReader(item.input))
			assert.NoError(t, err)
			assert.Equal(t, item.expect, result)
		}
	})

	t.Run("test parse array values", func(t *testing.T) {
		result, err := Parse(strings.NewReader("hello[1,2,3, 'four', five(start=1, end=2)]"))
		assert.NoError(t, err)
		first := result[0]
		assert.Equal(t, "hello", first.Name)
		assert.Equal(t, 0, first.Position)
		assert.Equal(t, 5, len(first.Array))
		assert.Equal(t, ptr("1"), first.Array[0].Value.Number)
		assert.Equal(t, ptr("2"), first.Array[1].Value.Number)
		assert.Equal(t, ptr("3"), first.Array[2].Value.Number)
		assert.Equal(t, ptr("four"), first.Array[3].Value.String)
		assert.Equal(t, "five", first.Array[4].Name)
		assert.Equal(t, "start", first.Array[4].Object[0].Name)
		assert.Equal(t, "end", first.Array[4].Object[1].Name)
		assert.Equal(t, ptr("1"), first.Array[4].Object[0].Value.Number)
		assert.Equal(t, ptr("2"), first.Array[4].Object[1].Value.Number)

		assert.Equal(t, first.Name, "hello")
	})

}
