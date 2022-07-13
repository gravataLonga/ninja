package stdlib

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
)

// Puts print stuff to standard output
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
