package stdlib

import "github.com/gravataLonga/ninja/object"

// Builtins register functions at global state.
// this is same as registering a function in Global Namespace
// E.g.: object.GlobalEnvironment.Set("rest", object.NewBuiltin(Rest))
var Builtins = map[string]*object.Builtin{}
