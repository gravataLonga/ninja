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

// Assign walks up the environment chain to find and update an existing binding.
// Returns false if the variable is not found anywhere in the chain.
func (e *Environment) Assign(name string, val Object) bool {
	if _, ok := e.store[name]; ok {
		e.store[name] = val
		return true
	}
	if e.outer != nil {
		return e.outer.Assign(name, val)
	}
	return false
}

func (e *Environment) ancestor(distance int) *Environment {
	env := e
	for i := 0; i < distance; i++ {
		env = env.outer
	}
	return env
}

func (e *Environment) GetAt(distance int, name string) Object {
	return e.ancestor(distance).store[name]
}

func (e *Environment) SetAt(distance int, name string, val Object) {
	e.ancestor(distance).store[name] = val
}
