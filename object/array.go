package object

import (
	"bytes"
	"strconv"
	"strings"
)

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := make([]string, len(ao.Elements))
	for i, e := range ao.Elements {
		elements[i] = e.Inspect()
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

func (s *Array) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method type not accept any arguments. got: %s", argStr)
		}
		return &String{Value: ARRAY_OBJ}
	case "join":
		return arrayJoin(s.Elements, args...)
	case "push":
		return arrayPush(s.Elements, args...)
	case "pop":
		return arrayPop(s.Elements, args...)
	case "shift":
		return arrayShift(s.Elements, args...)
	case "slice":
		return arraySlice(s.Elements, args...)
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}

func arrayJoin(elements []Object, args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("array.join expect exactly 1 argument. Got: %d", len(args))
	}

	joinArgument, ok := args[0].(*String)
	if !ok {
		return NewErrorFormat("array.join expect first argument be string. Got: %s", args[0].Type())
	}

	var out bytes.Buffer
	elementsString := []string{}
	for _, el := range elements {
		switch el.(type) {
		case *String:
			v, _ := el.(*String)
			elementsString = append(elementsString, v.Value)
		case *Integer:
			v, _ := el.(*Integer)
			elementsString = append(elementsString, strconv.Itoa(int(v.Value)))
		case *Float:
			v, _ := el.(*Float)
			elementsString = append(elementsString, strconv.FormatFloat(v.Value, 'f', -1, 64))
		case *Null:
			elementsString = append(elementsString, NULL_OBJ)
		case *Boolean:
			v, _ := el.(*Boolean)
			if v.Value {
				elementsString = append(elementsString, TRUE.Inspect())
			} else {
				elementsString = append(elementsString, FALSE.Inspect())
			}
		case *Array:
			v, _ := el.(*Array)
			joinObject := v.Call("join", args...)
			strJoinObject, ok := joinObject.(*String)
			if !ok {
				return NewErrorFormat("Unable to join array")
			}

			elementsString = append(elementsString, strJoinObject.Value)
		case *Hash:
			return NewErrorFormat("Hash cant be join")
		}

	}

	out.WriteString("[")
	out.WriteString(strings.Join(elementsString, joinArgument.Value))
	out.WriteString("]")
	return &String{Value: out.String()}
}

func arrayPush(elements []Object, args ...Object) Object {
	if len(args) <= 0 {
		return NewErrorFormat("array.push expect exactly 1 argument. Got: %d", len(args))
	}

	arr := &Array{Elements: elements}

	for _, v := range args {
		arr.Elements = append(arr.Elements, v)
	}

	return arr
}

func arrayPop(elements []Object, args ...Object) Object {
	if len(args) > 0 {
		return NewErrorFormat("array.pop expect exactly 0 argument. Got: %d", len(args))
	}

	if len(elements) <= 0 {
		return NULL
	}
	return elements[len(elements)-1]
}

func arrayShift(elements []Object, args ...Object) Object {
	if len(args) > 0 {
		return NewErrorFormat("array.shift expect exactly 0 argument. Got: %d", len(args))
	}

	if len(elements) <= 0 {
		return NULL
	}
	return elements[0]
}

func arraySlice(elements []Object, args ...Object) Object {
	if len(args) <= 0 || len(args) >= 3 {
		return NewErrorFormat("array.slice(start, offset) expected at least 1 argument and at max 2 arguments. Got: %s", InspectArguments(args...))
	}

	start, ok := args[0].(*Integer)
	if !ok {
		return NewErrorFormat("array.slice(start, offset) first argument must be integer. Got: %s", args[0].Inspect())
	}

	maxLength := int64(len(elements))
	offset := maxLength

	if len(args) >= 2 && args[1] != nil {
		offsetInteger, ok := args[1].(*Integer)
		if !ok {
			return NewErrorFormat("array.slice(start, offset) second argument must be integer. Got: %s", args[1].Inspect())
		}

		offset = offsetInteger.Value + start.Value
	}

	if offset <= start.Value {
		offset = start.Value
	}

	newElements := elements[start.Value:offset]

	return &Array{Elements: newElements}
}
