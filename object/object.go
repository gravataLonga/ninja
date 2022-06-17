package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

type BuiltinFunction func(args ...Object) Object

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"

	RETURN_VALUE_OBJ = "RETURN_VALUE"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"

	INTEGER_OBJ = "INTEGER"
	FLOAT_OBJ   = "FLOAT"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ  = "STRING"
	ARRAY_OBJ   = "ARRAY"
	HASH_OBJ    = "HASH"
)

func IsError(o Object) bool {
	return o != nil && o.Type() == ERROR_OBJ
}

func IsTruthy(o Object) bool {
	switch o {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
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
