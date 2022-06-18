package stdlib

import "ninja/object"

var Builtins = map[string]*object.Builtin{}

func init() {
	RegisterBuiltIn("len", Len)
	RegisterBuiltIn("first", First)
	RegisterBuiltIn("last", Last)
	RegisterBuiltIn("push", Push)
	RegisterBuiltIn("rest", Rest)
	RegisterBuiltIn("puts", Puts)
}

func RegisterBuiltIn(name string, function object.BuiltinFunction) {
	Builtins[name] = object.NewBuiltin(function)
}
