package semantic

import "github.com/gravataLonga/ninja/ast"

type Semantic struct {
	program ast.Node
}

func New(node ast.Node) *Semantic {
	return &Semantic{program: node}
}

func (s *Semantic) Analysis() ast.Node {
	return s.analysis(s.program)
}

func analysis(node ast.Node) ast.Node {
	switch node := node.(type) {
	case *ast.Program:
		for i, stmt := range node.Statements {
			node.Statements[i] = analysis(stmt)
		}
		return node
	}
	return node
}
