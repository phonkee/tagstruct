package main

import "github.com/yuin/stagparser"

type User struct {
	Name string `ts:"required,length(min=4,max=10)"`
}

func main() {
	user := &User{"bob"}
	definitions, err := stagparser.ParseStruct(user, "ts")
	_, _ = definitions, err
}
