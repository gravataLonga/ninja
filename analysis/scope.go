package analysis

type Scope struct {
	s map[string]bool
}

func NewScope() *Scope {
	return &Scope{make(map[string]bool)}
}

func (scope *Scope) Put(key string, defined bool) {
	scope.s[key] = defined
}

func (scope *Scope) Exists(key string) bool {
	_, ok := scope.s[key]
	return ok
}

func (scope *Scope) Get(key string) bool {
	if !scope.Exists(key) {
		return false
	}
	return scope.s[key]
}
