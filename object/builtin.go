package object

import "fmt"

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// GetBuiltinByName takes a name, iterates over our builtins slice and returns
// the appropriate builtin
func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}

// Builtins defines all of Monkey's built in functions
var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{"len", &Builtin{Fn: bLen}},
	{"first", &Builtin{Fn: bFirst}},
	{"puts", &Builtin{Fn: bPuts}},
	{"last", &Builtin{Fn: bLast}},
	{"rest", &Builtin{Fn: bRest}},
	{"push", &Builtin{Fn: bPush}},
}

func bLen(args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Integer{Value: int64(len(arg.Elements))}
	case *String:
		return &Integer{Value: int64(len(arg.Value))}
	default:
		return NewErrorFormat("argument to `len` not supported, got %s", args[0].Type())
	}
}

func bFirst(args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != ARRAY_OBJ {
		return NewErrorFormat("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func bPuts(args ...Object) Object {
	for _, arg := range args {
		if arg == nil {
			fmt.Println("Argument is nil")
			continue
		}
		fmt.Println(arg.Inspect())
	}

	return nil
}

func bLast(args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != ARRAY_OBJ {
		return NewErrorFormat("argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func bRest(args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != ARRAY_OBJ {
		return NewErrorFormat("argument to `rest` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &Array{Elements: newElements}
	}

	return NULL
}

func bPush(args ...Object) Object {
	if len(args) != 2 {
		return NewErrorFormat("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != ARRAY_OBJ {
		return NewErrorFormat("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)

	newElements := make([]Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &Array{Elements: newElements}
}
