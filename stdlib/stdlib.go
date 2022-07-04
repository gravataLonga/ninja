package stdlib

import "ninja/object"

var Builtins = map[string]*object.Builtin{
	"len":   object.NewBuiltin(Len),
	"first": object.NewBuiltin(First),
	"last":  object.NewBuiltin(Last),
	"push":  object.NewBuiltin(Push),
	"rest":  object.NewBuiltin(Rest),
	"puts":  object.NewBuiltin(Puts),
	"time":  object.NewBuiltin(Time),
	"rand":  object.NewBuiltin(Rand),
	"args":  object.NewBuiltin(Args),
	"exit":  object.NewBuiltin(Exit),
}
