package object

type Environment struct {
	isGlobal bool
	store    map[string]Object
	outer    *Environment
}

var GlobalEnvironment = NewGlobalEnvironment()

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.isGlobal = false
	env.outer = outer
	return env
}

func NewGlobalEnvironment() *Environment {
	env := NewEnvironment()
	env.isGlobal = true
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{isGlobal: false, store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

/*
func (e *Environment) Set(name string, val Object) Object {
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		e.outer.Set(name, val)
	} else {
		e.store[name] = val
	}

	return val
}
*/
