# tagstruct

Parse tag values defined by struct

# Example

First we need to define available keywords

```go
type Tag struct {
    DefaultFirst int `ts:"name=default"`
}
```

And then we can parse fields of struct

```go
type Struct struct {
    First int `ts:"default=42"`
}
```

then we need to first parse the tag structure, and then we can reuse it

```go
    parsed := New(Tag)
	
```
