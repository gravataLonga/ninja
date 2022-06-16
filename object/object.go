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
