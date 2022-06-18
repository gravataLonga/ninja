package stdlib

import (
	"fmt"
	"ninja/object"
)

func Puts(args ...object.Object) object.Object {
	for _, arg := range args {
		if arg == nil {
			fmt.Println("Argument is nil")
			continue
		}
		fmt.Println(arg.Inspect())
	}

	return nil
}
