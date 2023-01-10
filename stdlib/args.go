package stdlib

import (
	"github.com/gravataLonga/ninja/object"
)

func init() {
	object.GlobalEnvironment.Set("args", object.NewBuiltin(Args))
}

// Args will get everthing from argument when executed from cli
func Args(args ...object.Object) object.Object {
	err := object.Check(
		"args", args,
		object.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	elements := make([]object.Object, len(object.Arguments))
	for i, arg := range object.Arguments {
		elements[i] = &object.String{Value: arg}
	}
	return &object.Array{Elements: elements}
}
