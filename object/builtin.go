package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin" }

func NewBuiltin(function BuiltinFunction) *Builtin {
	return &Builtin{Fn: function}
}
