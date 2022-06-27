package object

import (
	"hash/fnv"
	"strconv"
	"strings"
)

type String struct {
	Value        string
	cacheHashKey uint64
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

func (s *String) HashKey() HashKey {
	if s.cacheHashKey <= 0 {
		h := fnv.New64a()
		h.Write([]byte(s.Value))
		s.cacheHashKey = h.Sum64()
	}

	return HashKey{Type: s.Type(), Value: s.cacheHashKey}
}

func (s *String) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method type not accept any arguments. got: %s", argStr)
		}
		return &String{Value: STRING_OBJ}
	case "split":
		return stringSplit(s.Value, args...)
	case "contain":
		return stringContains(s.Value, args...)
	case "index":
		return stringIndex(s.Value, args...)
	case "upper":
		return stringUpper(s.Value, args...)
	case "lower":
		return stringLower(s.Value, args...)
	case "trim":
		return stringTrim(s.Value, args...)
	case "int":
		return stringInteger(s.Value, args...)
	case "float":
		return stringFloat(s.Value, args...)
	}
	return NewErrorFormat("method %s not exists on string object.", method)
}

// @todo test for unhappy path for all methods
// @todo performance for all methods

func stringSplit(str string, args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("split method expect exactly 1 argument")
	}

	split, ok := args[0].(*String)
	if !ok {
		return NewErrorFormat("first argument must be string, got: %s", args[0].Type())
	}
	arr := strings.Split(str, split.Value)

	arrObject := &Array{}
	for _, i := range arr {
		arrObject.Elements = append(arrObject.Elements, &String{Value: i})
	}
	return arrObject
}

func stringContains(str string, args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("contain method expect exactly 1 argument")
	}

	needle, ok := args[0].(*String)
	if !ok {
		return NewErrorFormat("first argument must be string, got: %s", args[0].Type())
	}

	if strings.Contains(str, needle.Value) {
		return TRUE
	}
	return FALSE
}

func stringIndex(str string, args ...Object) Object {
	if len(args) != 1 {
		return NewErrorFormat("index method expect exactly 1 argument")
	}

	needle, ok := args[0].(*String)
	if !ok {
		return NewErrorFormat("first argument must be string, got: %s", args[0].Type())
	}

	val := strings.Index(str, needle.Value)
	return &Integer{Value: int64(val)}
}

func stringUpper(str string, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("upper method expect exactly 0 argument")
	}

	val := strings.ToUpper(str)
	return &String{Value: val}
}

func stringLower(str string, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("lower method expect exactly 0 argument")
	}
	val := strings.ToLower(str)
	return &String{Value: val}
}

func stringTrim(str string, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("trim method expect exactly 0 argument")
	}
	val := strings.Trim(str, "\n\r\t ")
	return &String{Value: val}
}

func stringInteger(str string, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("int method expect exactly 0 argument")
	}

	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return NewErrorFormat("string.int() fail to convert to int. Got: %s", err)
	}
	return &Integer{Value: val}
}

func stringFloat(str string, args ...Object) Object {
	if len(args) != 0 {
		return NewErrorFormat("float method expect exactly 0 argument")
	}

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return NewErrorFormat("string.float() fail to convert to float. Got: %s", err)
	}
	return &Float{Value: val}
}
