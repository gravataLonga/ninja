package object

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

// EPSILON is used only for internally usage only for now.
var EPSILON float64 = 0.00000001

type Float struct {
	Value        float64
	hashKeyCache uint64
}

func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }
func (f *Float) Type() ObjectType { return FloatObj }

func (f *Float) HashKey() HashKey {

	if f.hashKeyCache <= 0 {
		h := fnv.New64a()
		h.Write([]byte(fmt.Sprintf("%f", f.Value)))
		f.hashKeyCache = h.Sum64()
	}

	return HashKey{Type: f.Type(), Value: f.hashKeyCache}
}

func (f *Float) Compare(right Object) int8 {
	if obj, ok := right.(*Float); ok {
		switch {
		case f.Value < obj.Value:
			return -1
		case f.Value > obj.Value:
			return 1
		default:
			return 0
		}
	}
	return -1
}

func (f *Float) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		err := Check(
			"float.type",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}

		return &String{Value: FloatObj}
	case "string":
		err := Check(
			"float.string",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &String{Value: strconv.FormatFloat(f.Value, 'f', -1, 64)}
	case "abs":
		err := Check(
			"float.abs",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &Float{Value: math.Abs(f.Value)}
	case "round":
		err := Check(
			"float.round",
			args,
			ExactArgs(0),
		)

		if err != nil {
			return NewError(err.Error())
		}
		return &Float{Value: math.Round(f.Value)}
	}
	return NewErrorFormat("method %s not exists on integer object.", method)

}
