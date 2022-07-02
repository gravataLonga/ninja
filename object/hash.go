package object

import (
	"bytes"
	"fmt"
	"ninja/ast"
	"strings"
)

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (s *Hash) Call(objectCall *ast.ObjectCall, method string, env *Environment, args ...Object) Object {
	switch method {
	case "type":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method type not accept any arguments. got: %s", argStr)
		}
		return &String{Value: HASH_OBJ}
	case "keys":
		return hashKeys(s.Pairs, args...)
	case "has":
		return hashHas(s.Pairs, args...)
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}

func hashKeys(keys map[HashKey]HashPair, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("hash.keys() expect 0 arguments. Got: %s", InspectArguments(args...))
	}
	elements := make([]Object, len(keys))
	i := 0
	for _, pair := range keys {
		elements[i] = pair.Key
		i++
	}

	return &Array{Elements: elements}
}

func hashHas(keys map[HashKey]HashPair, args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("hash.has() expect at least 1 argument. got: %s", InspectArguments(args...))
	}

	hashable, ok := args[0].(Hashable)
	if !ok {
		return NewErrorFormat("hash.has() first argument isnt hashable. got: %s", InspectArguments(args...))
	}

	_, ok = keys[hashable.HashKey()]
	if ok {
		return TRUE
	}
	return FALSE
}
