package stdlib

import (
	"github.com/gravataLonga/ninja/object"
	"time"
)

func init() {
	object.GlobalEnvironment.Set("time", object.NewBuiltin(Time))
}

// Time we get time in seconds
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
