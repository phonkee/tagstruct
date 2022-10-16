package tagstruct

import (
	"fmt"
	"reflect"

	"github.com/phonkee/go-collection"
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
		attribs := d.Attributes()
		if value, ok := attribs[s.Alias]; ok {
			if value, ok := value.(int64); ok {
				f := val.FieldByName(s.Name)
				if f.CanSet() {
					f.SetInt(value)
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
	val := reflect.ValueOf(into)
	if val.CanSet() {
		return nil, fmt.Errorf("cannot address value: %T", into)
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	coll := collection.Collection[stagparser.Definition](defs)
	// ignore for now
	result := coll.Filter(func(d stagparser.Definition) bool {
		return d.Name() != s.Alias
	})

	value := make([]string, 0)
	coll.Filter(func(d stagparser.Definition) bool {
		return d.Name() == s.Alias
	}).Each(func(d stagparser.Definition) {
		if v, ok := d.Attributes()[s.Alias]; ok {
			if v, ok := v.([]interface{}); ok {
				for _, s := range v {
					if s, ok := s.(string); ok {
						value = append(value, s)
					}
				}
			}
		}
	})
	f := val.FieldByName(s.Name)
	if f.CanSet() {
		f.Set(reflect.ValueOf(value))
	}

	return result, nil
}

type BoolProperty struct {
	BaseProperty
}

func (b *BoolProperty) Unmarshall(defs []stagparser.Definition, into interface{}) ([]stagparser.Definition, error) {
	val := reflect.ValueOf(into)
	if val.CanSet() {
		return nil, fmt.Errorf("cannot address value: %T", into)
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	result := make([]stagparser.Definition, 0, len(defs))
	for _, d := range defs {
		if d.Name() != b.Alias {
			result = append(result, d)
			continue
		}
		attribs := d.Attributes()
		if len(attribs) == 0 {
			f := val.FieldByName(b.Name)
			if f.CanSet() {
				f.SetBool(true)
				delete(attribs, b.Alias)
			}
		} else {
			if value, ok := attribs[b.Alias]; ok {
				if value, ok := value.(bool); ok {
					f := val.FieldByName(b.Name)
					if f.CanSet() {
						f.SetBool(value)
						delete(attribs, b.Alias)
					}
				} else {
					// TODO: try string

					return nil, fmt.Errorf("invalid type for %s: %T", b.Alias, value)
				}
			}
		}

	}

	return result, nil

}
