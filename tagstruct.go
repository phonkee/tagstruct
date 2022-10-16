package tagstruct

import (
	"fmt"
	"reflect"

	"github.com/yuin/stagparser"
)

// New analyzes given struct and returns definition. definition can then parse tags and returns values
// If something fails, this function panics
func New[T any](what T) Definition[T] {
	var props []Property

	// now we go over all fields and check which are used
	typ := reflect.TypeOf(what)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		// if field is not ignored, we need to check if it is defined in stag parser
		d, err := stagparser.ParseTag(field.Tag.Get("ts"), "ts")
		_ = d
		if err != nil {
			panic(err)
		}

		// now prepare base property
		base := BaseProperty{
			Name:  field.Name,
			Alias: field.Name,
		}
		for _, def := range d {
			if def.Name() == "name" {
				if name, ok := def.Attribute("name"); ok {
					if name, ok := name.(string); ok {
						base.Alias = name
					}
				}
			}
		}

		switch fieldType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			props = append(props, &IntProperty{
				BaseProperty: base,
			})
		case reflect.String:
			props = append(props, &StringProperty{
				BaseProperty: base,
			})
		case reflect.Array, reflect.Slice:
			props = append(props, &StringArrayProperty{
				BaseProperty: base,
			})
		default:
			panic(fmt.Sprintf("unsupported type %s", fieldType.Kind()))
		}
	}

	return Definition[T]{
		props: props,
	}
}
