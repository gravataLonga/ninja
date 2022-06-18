package stdlib

import "ninja/object"

func Len(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewErrorFormat("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return object.NewErrorFormat("argument to `len` not supported, got %s", args[0].Type())
	}
}
