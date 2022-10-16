package tagstruct

import "github.com/yuin/stagparser"

type Property interface {
	Unmarshall(what []stagparser.Definition, into interface{}) ([]stagparser.Definition, error)
}
