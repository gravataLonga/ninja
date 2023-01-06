package stdlib

import (
	"github.com/gravataLonga/ninja/object"
)

// Rest we get last items from array without first item
func Rest(args ...object.Object) object.Object {

	err := object.Check(
		"first", args,
		object.ExactArgs(1),
		object.WithTypes(object.ArrayObj),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return object.NULL
}
