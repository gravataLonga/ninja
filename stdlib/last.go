package stdlib

import (
	"github.com/gravataLonga/ninja/object"
)

func init() {
	object.GlobalEnvironment.Set("last", object.NewBuiltin(Last))
}

// Last get item from array
func Last(args ...object.Object) object.Object {

	err := object.Check(
		"last", args,
		object.ExactArgs(1),
		object.WithTypes(object.ARRAY_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	cloneable, ok := args[0].(object.Cloneable)
	if !ok {
		return object.NewErrorFormat("object isn't cloneable.")
	}

	arrClone := cloneable.Clone()
	arr := arrClone.(*object.Array)

	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return object.NULL
}
