package main

import "github.com/gravataLonga/ninja/object"

func Hello(args ...object.Object) object.Object {
	return &object.String{Value: "Hello World!"}
}

func main() {
	panic("this is a plugin")
	// Build a plugin: go build -buildmode=plugin hello.go
}
