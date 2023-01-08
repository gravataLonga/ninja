package stdlib

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
)

func init() {
	object.GlobalEnvironment.Set("puts", object.NewBuiltin(Puts))
}

// Puts print stuff to standard output
func Puts(args ...object.Object) object.Object {
	for _, arg := range args {
		if arg == nil {
			_, err := fmt.Fprintln(object.StandardOutput, "Argument is nil")
			if err != nil {
				return object.NewError("Unable to put to standard output")

			}
			continue
		}
		fmt.Println(arg.Inspect())
	}

	return nil
}
