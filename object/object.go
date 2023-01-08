package object

import (
	"bytes"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// CallableMethod is use for calling method in data type.
type CallableMethod interface {
	Call(method string, args ...Object) Object
}

type Cloneable interface {
	Clone() Object
}

// Comparable is the interface for comparing two Object and their underlying
// values. It is the responsibility of the caller (left) to check for types.
// E.g.: 1 > 1, it will return -1. left isn't greater than 1.
type Comparable interface {
	Compare(right Object) int8
}

// HashKey hold "key" on Hash
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// Hashable exist for object implement it in order to be used in Hash object
type Hashable interface {
	HashKey() HashKey
}

const (
	NULL_OBJ         = "NULL"
	ERROR_OBJ        = "ERROR"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	BREAK_VALUE_OBJ  = "BREAK_VALUE"
	ENUM_OBJ         = "ENUM"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	PLUGIN_OBJ       = "PLUGIN"
)

func IsError(o Object) bool {
	return o != nil && o.Type() == ERROR_OBJ
}

func IsTruthy(o Object) bool {
	if o == nil {
		return false
	}

	switch o.Type() {
	case NULL_OBJ:
		return false
	case BOOLEAN_OBJ:
		v, ok := o.(*Boolean)
		if !ok {
			return false
		}
		return v.Value
	default:
		return true
	}
}

func IsNumber(o Object) bool {
	return o != nil && (o.Type() == INTEGER_OBJ || o.Type() == FLOAT_OBJ)
}

func IsArray(o Object) bool {
	return o != nil && o.Type() == ARRAY_OBJ
}

func IsHash(o Object) bool {
	return o != nil && o.Type() == HASH_OBJ
}

func IsString(o Object) bool {
	return o != nil && o.Type() == STRING_OBJ
}

func InspectObject(args ...Object) string {
	var out bytes.Buffer
	elements := make([]string, len(args))
	for i, e := range args {
		elements[i] = e.Inspect()
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
