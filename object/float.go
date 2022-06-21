package object

import (
	"fmt"
	"hash/fnv"
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
