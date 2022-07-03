package stdlib

import (
	"ninja/object"
	"time"
)

func Time(args ...object.Object) object.Object {
	err := object.Check(
		"first", args,
		object.ExactArgs(0),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	return &object.Integer{Value: time.Now().Unix()}
}
