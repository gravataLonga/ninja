package object

import (
	"hash/fnv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type String struct {
	Value         string
	hashKeyCached uint64
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

func (s *String) HashKey() HashKey {
	if s.hashKeyCached == 0 {
		h := fnv.New64a()
		h.Write([]byte(s.Value))
		s.hashKeyCached = h.Sum64()
	}

	return HashKey{Type: s.Type(), Value: s.hashKeyCached}
}

func (s *String) Compare(right Object) int8 {
	str, ok := right.(*String)
	if !ok {
		return -1
	}

	if str.Value == s.Value {
		return 1
	}

	return 0
}

func (s *String) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"string.type",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}

		return &String{Value: STRING_OBJ}
	case "length":
		err := Check(
			"string.length",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
	case "split":
		return stringSplit(s.Value, args...)
	case "replace":
		return stringReplace(s.Value, args...)
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

// @todo performance for all methods

func stringSplit(str string, args ...Object) Object {
	err := Check(
		"string.split",
		args,
		ExactArgs(1),
		WithTypes(STRING_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

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

func stringReplace(str string, args ...Object) Object {
	err := Check(
		"string.replace",
		args,
		ExactArgs(2),
		WithTypes(STRING_OBJ, STRING_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

	search, _ := args[0].(*String)
	replace, _ := args[1].(*String)

	newStr := strings.ReplaceAll(str, search.Value, replace.Value)
	return &String{Value: newStr}
}

func stringContains(str string, args ...Object) Object {

	err := Check(
		"string.contain",
		args,
		ExactArgs(1),
		WithTypes(STRING_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

	needle, _ := args[0].(*String)

	if strings.Contains(str, needle.Value) {
		return TRUE
	}
	return FALSE
}

func stringIndex(str string, args ...Object) Object {
	err := Check(
		"string.index",
		args,
		ExactArgs(1),
		WithTypes(STRING_OBJ),
	)

	if err != nil {
		return NewError(err.Error())
	}

	needle, _ := args[0].(*String)

	val := strings.Index(str, needle.Value)
	return &Integer{Value: int64(val)}
}

func stringUpper(str string, args ...Object) Object {
	err := Check(
		"string.upper",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	val := strings.ToUpper(str)
	return &String{Value: val}
}

func stringLower(str string, args ...Object) Object {
	err := Check(
		"string.lower",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	val := strings.ToLower(str)
	return &String{Value: val}
}

func stringTrim(str string, args ...Object) Object {
	err := Check(
		"string.trim",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}
	val := strings.Trim(str, "\n\r\t ")
	return &String{Value: val}
}

func stringInteger(str string, args ...Object) Object {
	err := Check(
		"string.int",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return NewErrorFormat("TypeError: string.int() fail to convert to int. Got: %s", err)
	}
	return &Integer{Value: val}
}

func stringFloat(str string, args ...Object) Object {
	err := Check(
		"string.float",
		args,
		ExactArgs(0),
	)

	if err != nil {
		return NewError(err.Error())
	}

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return NewErrorFormat("TypeError: string.float() fail to convert to float. Got: %s", err)
	}
	return &Float{Value: val}
}
