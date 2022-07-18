package object

import (
	"fmt"
	"strings"
)

// CheckFunc signature for check arguments function
type CheckFunc func(name string, args []Object) error

// Check arguments passed to object call we can pass many CheckFunc as you want
func Check(name string, args []Object, checks ...CheckFunc) error {
	for _, check := range checks {
		if err := check(name, args); err != nil {
			return err
		}
	}
	return nil
}

// ExactArgs expect exact nums arguments
func ExactArgs(n int) CheckFunc {
	return func(name string, args []Object) error {
		if len(args) != n {
			return fmt.Errorf(
				"TypeError: %s() takes exactly %d argument (%d given)",
				name, n, len(args),
			)
		}
		return nil
	}
}

// MinimumArgs check if have at least n arguments
func MinimumArgs(n int) CheckFunc {
	return func(name string, args []Object) error {
		if len(args) < n {
			return fmt.Errorf(
				"TypeError: %s() takes a minimum %d arguments (%d given)",
				name, n, len(args),
			)
		}
		return nil
	}
}

// RangeOfArgs expect n until m arguments
func RangeOfArgs(n, m int) CheckFunc {
	return func(name string, args []Object) error {
		if len(args) < n || len(args) > m {
			return fmt.Errorf(
				"TypeError: %s() takes at least %d arguments at most %d (%d given)",
				name, n, m, len(args),
			)
		}
		return nil
	}
}

// WithTypes combined with ExactArgs it will check if we got ObjectType by it is order.
func WithTypes(types ...ObjectType) CheckFunc {
	return func(name string, args []Object) error {
		for i, t := range types {
			if i < len(args) && args[i].Type() != t {
				return fmt.Errorf(
					"TypeError: %s() expected argument #%d to be `%s` got `%s`",
					name, (i + 1), t, args[i].Type(),
				)
			}
		}
		return nil
	}
}

// OneOfType almost same as WithTypes but check one of type of ObjectType
func OneOfType(types ...ObjectType) CheckFunc {
	return func(name string, args []Object) error {
		for _, vt := range types {
			for _, va := range args {
				if va.Type() == vt {
					return nil
				}
			}
		}

		expected := make([]string, len(types))
		for i, t := range types {
			expected[i] = string(t)
		}

		got := make([]string, len(args))
		for i, t := range args {
			got[i] = string(t.Type())
		}

		return fmt.Errorf(
			"TypeError: %s() expected argument to be `%s` got `%s`",
			name, strings.Join(expected, ","), strings.Join(got, ","),
		)
	}
}
