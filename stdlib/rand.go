package stdlib

import (
	"math/rand"
	"ninja/object"
)

// Rand generate a random number from 0 ... 1 float
func Rand(args ...object.Object) object.Object {
	err := object.Check(
		"rand", args,
		object.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	return &object.Float{Value: rand.Float64()}
}
