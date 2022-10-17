# tagstruct

This package aims to provide a simple and efficient way to define struct tag parsers.
Currently it's very limited, but it's a start.
We can define parser that parses e.g. this:

    value="string", int=1, bool=true, float32=1.0, float64=2.0

# Example

First we need to define available keywords

```go
type Tag struct {
    DefaultFirst int `ts:"name=default"`
}
```

And then we can parse fields of struct

```go
tagdef := New(Tag)
result, err := tagdef.Parse("default=42")
```

# TODO:

We like to be able to parse objects (structs) in recursive manner and also arrays, such as:

    `some(id=1, value=2, span(from=1, to=2))`

Even arrays

    `some[(id=1), (id=2)]`

For this we will need to roll our own parser, but it's not that hard.
