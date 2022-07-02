package stdlib

import (
	"math/rand"
	"ninja/object"
	"ninja/typing"
)

func Rand(args ...object.Object) object.Object {
	err := typing.Check(
		"rand", args,
		typing.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	return &object.Float{Value: rand.Float64()}
}
