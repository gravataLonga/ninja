package object

import "hash/fnv"

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
