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

type Closure interface {
	Call(args ...Object) Object
}

const (
	NullObj        = "NULL"
	ErrorObj       = "ERROR"
	ReturnValueObj = "RETURN_VALUE"
	BreakValueObj  = "BREAK_VALUE"
	EnumObj        = "ENUM"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
	IntegerObj     = "INTEGER"
	FloatObj       = "FLOAT"
	BooleanObj     = "BOOLEAN"
	StringObj      = "STRING"
	ArrayObj       = "ARRAY"
	HashObj        = "HASH"
)

func IsError(o Object) bool {
	return o != nil && o.Type() == ErrorObj
}

func IsTruthy(o Object) bool {
	if o == nil {
		return false
	}

	switch o.Type() {
	case NullObj:
		return false
	case BooleanObj:
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
	return o != nil && (o.Type() == IntegerObj || o.Type() == FloatObj)
}

func IsArray(o Object) bool {
	return o != nil && o.Type() == ArrayObj
}

func IsHash(o Object) bool {
	return o != nil && o.Type() == HashObj
}

func IsString(o Object) bool {
	return o != nil && o.Type() == StringObj
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
