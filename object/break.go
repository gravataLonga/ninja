package object

type Break struct {
	Value Object
}

func (b *Break) Type() ObjectType { return BreakValueObj }
func (b *Break) Inspect() string  { return b.Value.Inspect() }
