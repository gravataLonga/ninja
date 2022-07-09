package semantic

import (
	"fmt"
	"ninja/ast"
)

// Semantic here is where we are going doing some semantic analysis
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

func (s *Semantic) Analysis(node *ast.Program) ast.Node {
	return s.analysis(node)
}

// newScope record scope how deep is
func (s *Semantic) newScope() {
	s.scopeStack.Push(&Scope{})
}

// exitScope remove one scope top of head of scope
func (s *Semantic) exitScope() {
	s.scopeStack.Pop()
}

// declare will keep track of declare variables
func (s *Semantic) declare(node *ast.VarStatement) {
	if s.scopeStack.IsEmpty() {
		return
	}

	peek, _ := s.scopeStack.Peek()
	(*peek)[node.Name.Value] = false
}

// resolve after a variable been resolve we mark it as resolved.
func (s *Semantic) resolve(node *ast.VarStatement) {
	name := node.Name.Value
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

func (s *Semantic) analysis(node ast.Node) ast.Node {
	switch node := node.(type) {
	case *ast.Program:
		for _, v := range node.Statements {
			s.analysis(v)
		}
	case *ast.ArrayLiteral:
		for _, e := range node.Elements {
			s.analysis(e)
		}
	case *ast.IfExpression:
		s.analysis(node.Condition)
		s.analysis(node.Consequence)
		s.analysis(node.Alternative)
	case *ast.HashLiteral:
		for k, v := range node.Pairs {
			s.analysis(k)
			s.analysis(v)
		}
	case *ast.Identifier:
		s.expectIdentifierDeclare(node)
	case *ast.VarStatement:
		s.declare(node)
		if _, ok := node.Value.(ast.Expression); ok {
			s.analysis(node.Value)
		}
		s.resolve(node)
	case *ast.ExpressionStatement:
		s.analysis(node.Expression)
	case *ast.FunctionLiteral:
		s.analysis(node.Body)
	case *ast.BlockStatement:
		s.newScope()
		for _, b := range node.Statements {
			s.analysis(b)
		}
		s.exitScope()
	}
	return node
}
