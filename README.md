# tagstruct

Parse tag values defined by struct

# Example

Let's suppose we want to be able to parse these attributes

```go
type Example struct {
    Field int `custom:"rename=something"`
}
```
