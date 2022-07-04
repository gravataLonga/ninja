package stdlib

import "ninja/object"

// ExitFunction

// Exit execute exit function. Terminate following program
func Exit(args ...object.Object) object.Object {

	err := object.Check(
		"exit", args,
		object.ExactArgs(1),
		object.WithTypes(object.INTEGER_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	intV, _ := args[0].(*object.Integer)

	object.ExitFunction(int(intV.Value))
	return object.NULL
}
