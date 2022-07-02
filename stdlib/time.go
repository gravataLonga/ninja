package stdlib

import (
	"ninja/object"
	"ninja/typing"
	"time"
)

func Time(args ...object.Object) object.Object {
	err := typing.Check(
		"first", args,
		typing.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	return &object.Integer{Value: time.Now().Unix()}
}
