package parser

// Property represents property in tag
type Property struct {
	// name of property
	Name string

	// position in original string
	Position int

	// This is poor man's union in go (not complaining, just saying)
	Value  *Value
	Object []Property
	Array  []Property
}

// HasValue returns whether any value was set to property
func (p *Property) HasValue() bool {
	return p.Value != nil || p.Object != nil || p.Array != nil
}

type Value struct {
	Position int

	String *string
	Number *string
}
