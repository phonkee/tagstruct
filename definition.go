package tagstruct

import (
	"fmt"
	"reflect"

	"github.com/yuin/stagparser"
)

type Definition[T any] struct {
	props []Property
}

// ParseStruct parses all struct fields
func (d Definition[T]) ParseStruct(what interface{}, tag string) (map[string]T, error) {
	val := reflect.ValueOf(what)
	typ := reflect.TypeOf(what)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", val.Kind())
	}

	result := map[string]T{}

	for i := 0; i < val.NumField(); i++ {
		p, err := d.ParseTag(typ.Field(i).Tag.Get(tag))
		if err != nil {
			return nil, err
		}
		result[typ.Field(i).Name] = p
	}

	return result, nil
}

// ParseTag parses struct field tag and returns struct with all information
func (d Definition[T]) ParseTag(tag string) (T, error) {
	result := new(T)
	defs, err := stagparser.ParseTag(tag, "")
	if err != nil {
		return *result, err
	}

	for _, prop := range d.props {
		if defs, err = prop.Unmarshall(defs, result); err != nil {
			return *result, err
		}
	}

	if len(defs) > 0 {
		names := make([]string, 0)
		for _, d := range defs {
			names = append(names, d.Name())
		}
		return *result, fmt.Errorf("unknown fields %v", names)
	}

	return *result, nil
}
