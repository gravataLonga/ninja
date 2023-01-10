package main

import . "github.com/gravataLonga/ninja/object"

func Hello(args ...Object) Object {
	return &String{Value: "Hello World!"}
}

func main() {
	panic("this is a plugin")
	// Build a plugin:go build -buildmode=plugin -o fixtures/hello.so fixtures/hello.go
}
