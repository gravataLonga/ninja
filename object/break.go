package object

type Break struct {
	Value Object
}

func (b *Break) Type() ObjectType { return BREAK_VALUE_OBJ }
func (b *Break) Inspect() string  { return b.Value.Inspect() }
