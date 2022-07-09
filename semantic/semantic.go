package semantic

import (
	"fmt"
	"ninja/ast"
)

// Semantic here is where we are going doing some semantic analyze
// using visitor pattern
type Semantic struct {
	scopeStack     Stack
	globalVariable map[string]ast.Expression
	errors         []string
}

func New() *Semantic {
	return &Semantic{}
}

func (s *Semantic) Errors() []string {
	return s.errors
}

func (s *Semantic) NewError(format string, a ...interface{}) {
	s.errors = append(s.errors, fmt.Sprintf(format, a...))
}

func (s *Semantic) Analyze(node *ast.Program) ast.Node {
	return s.analyze(node)
}

// beginScope record scope how deep is
func (s *Semantic) beginScope() {
	s.scopeStack.Push(&Scope{})
}

// endScope remove one scope top of head of scope
func (s *Semantic) endScope() {
	s.scopeStack.Pop()
}

// declare will keep track of declare variables
func (s *Semantic) declare(node *ast.VarStatement) {
	peek, ok := s.scopeStack.Peek()
	if !ok {
		return
	}

	name := node.Name.Value
	*peek = Scope{name: false}
}

// resolve after a variable been resolve we mark it as resolved.
func (s *Semantic) resolve(node *ast.VarStatement) {
	name := node.Name.Value
	if !s.expectIdentifierDeclare(node.Name) {
		return
	}

	peek, ok := s.scopeStack.Peek()
	if !ok {
		return
	}

	*peek = Scope{name: true}
}

func (s *Semantic) expectIdentifierDeclare(node *ast.Identifier) bool {
	name := node.Value
	peek, ok := s.scopeStack.Peek()
	if !ok {
		s.NewError("There aren't any scope active %s", name)
		return false
	}

	v, ok := (*peek)[name]
	if !ok {
		s.NewError("Variable \"%s\" not declare yet %s", name, node.Token)
		return false
	}

	if !v {
		s.NewError("Can't read local variable \"%s\" in its own initializer %s", name, node.Token)
		return false
	}

	return true
}

func (s *Semantic) analyze(node ast.Node) ast.Node {
	switch node := node.(type) {
	case *ast.Program:
		s.beginScope()
		for _, v := range node.Statements {
			s.analyze(v)
		}
		s.endScope()
	case *ast.ArrayLiteral:
		for _, e := range node.Elements {
			s.analyze(e)
		}
	case *ast.IfExpression:
		s.analyze(node.Condition)
		s.analyze(node.Consequence)
		s.analyze(node.Alternative)
	case *ast.HashLiteral:
		for k, v := range node.Pairs {
			s.analyze(k)
			s.analyze(v)
		}
	case *ast.Identifier:
		s.expectIdentifierDeclare(node)
	case *ast.VarStatement:
		s.declare(node)
		s.analyze(node.Value)
		s.resolve(node)
	case *ast.ExpressionStatement:
		s.analyze(node.Expression)
	case *ast.FunctionLiteral:
		s.analyze(node.Body)
	case *ast.BlockStatement:
		s.beginScope()
		for _, b := range node.Statements {
			s.analyze(b)
		}
		s.endScope()
	}
	return node
}
