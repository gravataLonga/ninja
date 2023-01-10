package stdlib

import "github.com/gravataLonga/ninja/object"

func init() {
	object.GlobalEnvironment.Set("plugin", object.NewBuiltin(Plugin))
}

// Plugin will load plugin into global environment
func Plugin(args ...object.Object) object.Object {

	err := object.Check(
		"plugin", args,
		object.ExactArgs(1),
		object.WithTypes(object.STRING_OBJ),
	)

	if err != nil {
		return object.NewError(err.Error())
	}

	str, _ := args[0].(*object.String)

	plugin, err := object.NewPlugin(str.Value)
	if err != nil {
		return object.NewError(err.Error())
	}
	return plugin
}
