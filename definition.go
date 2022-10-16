package tagstruct

import (
	"fmt"

	"github.com/yuin/stagparser"
)

type Definition[T any] struct {
	props []Property
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
		return *result, fmt.Errorf("unknown fields %#v", defs)
	}

	return *result, nil
}
