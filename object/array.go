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

func (s *Array) Clone() Object {
	arr := Array{}
	elems := make([]Object, len(s.Elements))
	for i, e := range s.Elements {
		cloneable, ok := e.(Cloneable)
		if ok {
			elems[i] = cloneable.Clone()
		} else {
			elems[i] = e
		}

	}
	arr.Elements = elems
	return &arr
}

func (s *Array) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"array.type", args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}

		return &String{Value: ARRAY_OBJ}
	case "length":
		err := Check(
			"array.length", args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}

		return &Integer{Value: int64(len(s.Elements))}
	case "join":
		return arrayJoin(s.Elements, args...)
	case "push":
		result := arrayPush(s, args...)
		return result
	case "pop":
		popValue := arrayPop(s, args...)
		return popValue
	case "shift":
		shiftValue := arrayShift(s, args...)
		return shiftValue
	case "slice":
		return arraySlice(s.Elements, args...)
	}
	return NewErrorFormat("method %s not exists on array object.", method)
}

func arrayJoin(elements []Object, args ...Object) Object {
	err := Check(
		"array.join", args,
		ExactArgs(1),
		WithTypes(STRING_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

	joinArgument, _ := args[0].(*String)

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

	out.WriteString(strings.Join(elementsString, joinArgument.Value))
	return &String{Value: out.String()}
}

func arrayPush(array *Array, args ...Object) Object {
	err := Check(
		"array.push", args,
		MinimumArgs(1),
	)

	if err != nil {
		return NewError(err.Error())
	}

	for _, v := range args {
		array.Elements = append(array.Elements, v)
	}

	return NULL
}

func arrayPop(array *Array, args ...Object) Object {
	err := Check(
		"array.pop", args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	if len(array.Elements) <= 0 {
		return NULL
	}

	poppedElement := array.Elements[len(array.Elements)-1]
	array.Elements = array.Elements[:len(array.Elements)-1]

	return poppedElement
}

func arrayShift(array *Array, args ...Object) Object {
	err := Check(
		"array.shift", args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	if len(array.Elements) <= 0 {
		return NULL
	}

	shiftedValue := array.Elements[0]
	array.Elements = array.Elements[1:]
	return shiftedValue
}

func arraySlice(elements []Object, args ...Object) Object {
	err := Check(
		"array.push", args,
		RangeOfArgs(1, 2),
		WithTypes(INTEGER_OBJ, INTEGER_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

	start, _ := args[0].(*Integer)

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
