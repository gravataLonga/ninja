package object

import (
	"bytes"
	"fmt"
	"strings"
)

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HashObj }
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

func (s *Hash) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"hash.type",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}

		return &String{Value: HashObj}
	case "keys":
		return hashKeys(s.Pairs, args...)
	case "values":
		return hashValues(s.Pairs, args...)
	case "has":
		return hashHas(s.Pairs, args...)
	case "merge":
		return hashMerge(s.Pairs, args...)
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}

func hashKeys(keys map[HashKey]HashPair, args ...Object) Object {
	err := Check(
		"hash.keys",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	elements := make([]Object, len(keys))
	i := 0
	for _, pair := range keys {
		elements[i] = pair.Key
		i++
	}

	return &Array{Elements: elements}
}

func hashValues(keys map[HashKey]HashPair, args ...Object) Object {
	err := Check(
		"hash.values",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	elements := make([]Object, len(keys))
	i := 0
	for _, pair := range keys {
		elements[i] = pair.Value
		i++
	}

	return &Array{Elements: elements}
}

func hashMerge(pairs map[HashKey]HashPair, args ...Object) Object {
	err := Check(
		"hash.merge",
		args,
		ExactArgs(1),
		WithTypes(HashObj),
	)

	if err != nil {
		return NewError(err.Error())
	}

	hashPairArg, _ := args[0].(*Hash)

	for k, v := range pairs {
		hashPairArg.Pairs[k] = v
	}

	return hashPairArg
}

func hashHas(keys map[HashKey]HashPair, args ...Object) Object {
	err := Check(
		"hash.has",
		args,
		ExactArgs(1),
	)

	if err != nil {
		return NewError(err.Error())
	}

	hashable, ok := args[0].(Hashable)
	if !ok {
		return NewErrorFormat("hash.has() first argument isnt hashable. got: %s", InspectObject(args...))
	}

	_, ok = keys[hashable.HashKey()]
	if ok {
		return TRUE
	}
	return FALSE
}
