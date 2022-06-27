package object

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

var EPSILON float64 = 0.00000001

type Float struct {
	Value        float64
	cacheHashKey uint64
}

func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }

func (f *Float) HashKey() HashKey {

	if f.cacheHashKey <= 0 {
		h := fnv.New64a()
		h.Write([]byte(fmt.Sprintf("%f", f.Value)))
		f.cacheHashKey = h.Sum64()
	}

	return HashKey{Type: f.Type(), Value: f.cacheHashKey}
}

func (f *Float) Call(method string, args ...Object) Object {
	switch method {
	case "type":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method type not accept any arguments. got: %s", argStr)
		}
		return &String{Value: FLOAT_OBJ}
	case "string":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method string not accept any arguments. got: %s", argStr)
		}
		return &String{Value: strconv.FormatFloat(f.Value, 'f', -1, 64)}
	case "abs":
		if len(args) > 0 {
			argStr := InspectArguments(args...)
			return NewErrorFormat("method abs not accept any arguments. got: %s", argStr)
		}
		var absT float64 = math.Abs(f.Value)
		return &Float{Value: absT}
	}
	return NewErrorFormat("method %s not exists on integer object.", method)

}
