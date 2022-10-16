package tagstruct

import (
	"fmt"
	"reflect"

	"github.com/yuin/stagparser"
)

type BaseProperty struct {
	Name  string
	Alias string
}

type IntProperty struct {
	BaseProperty
}

func (s *IntProperty) Unmarshall(defs []stagparser.Definition, into interface{}) ([]stagparser.Definition, error) {
	val := reflect.ValueOf(into)
	if val.CanSet() {
		return nil, fmt.Errorf("cannot address value: %T", into)
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	result := make([]stagparser.Definition, 0, len(defs))
	for _, d := range defs {
		if d.Name() != s.Alias {
			result = append(result, d)
			continue
		}
		result = append(result, d)
		attribs := d.Attributes()
		if value, ok := attribs[s.Alias]; ok {
			if value, ok := value.(int64); ok {
				f := val.FieldByName(s.Name)
				if f.CanSet() {
					f.SetInt(int64(value))
					delete(attribs, s.Alias)
				}
			} else {
				return nil, fmt.Errorf("invalid type for %s: %T", s.Alias, value)
			}
		}
	}

	return result, nil
}

type StringProperty struct {
	BaseProperty
}

func (s *StringProperty) Unmarshall(defs []stagparser.Definition, into interface{}) ([]stagparser.Definition, error) {
	val := reflect.ValueOf(into)
	if val.CanSet() {
		return nil, fmt.Errorf("cannot address value: %T", into)
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	result := make([]stagparser.Definition, 0, len(defs))
	for _, d := range defs {
		if d.Name() != s.Alias {
			result = append(result, d)
			continue
		}
		attribs := d.Attributes()
		if value, ok := attribs[s.Alias]; ok {
			if value, ok := value.(string); ok {
				f := val.FieldByName(s.Name)
				if f.CanSet() {
					f.SetString(value)
					delete(attribs, s.Alias)
				}
			} else {
				return nil, fmt.Errorf("invalid type for %s: %T", s.Alias, value)
			}
		}
	}

	return result, nil
}

type StringArrayProperty struct {
	BaseProperty
}

func (s *StringArrayProperty) Unmarshall(defs []stagparser.Definition, into interface{}) ([]stagparser.Definition, error) {
	// ignore for now
	return nil, nil
}
