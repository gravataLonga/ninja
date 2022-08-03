package semantic

import "github.com/gravataLonga/ninja/ast"

type Semantic struct {
	program ast.Node
}

func New(node ast.Node) *Semantic {
	return &Semantic{program: node}
}

func (s *Semantic) Analysis() ast.Node {
	return analysis(s.program)
}

func analysis(node ast.Node) ast.Node {
	return node
}
