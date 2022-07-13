package stdlib

import (
	"github.com/gravataLonga/ninja/object"
)

// Len will get length of object
func Len(args ...object.Object) object.Object {
	err := object.Check(
		"len", args,
		object.ExactArgs(1),
		object.OneOfType(object.ARRAY_OBJ, object.STRING_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
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
