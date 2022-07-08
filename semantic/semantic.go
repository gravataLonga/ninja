package semantic

import (
	"fmt"
	"ninja/ast"
	"ninja/parser"
)

// Semantic here is where we are going doing some semantic analyze
// using visitor pattern
type Semantic struct {
	p          *parser.Parser
	scopeStack Stack
	errors     []string
}

func New(p *parser.Parser) *Semantic {
	return &Semantic{p: p}
}

func (s *Semantic) Errors() []string {
	return s.errors
}

func (s *Semantic) NewError(format string, a ...interface{}) {
	s.errors = append(s.errors, fmt.Sprintf(format, a...))
}

func (s *Semantic) Analyze() ast.Node {
	return s.analyze(s.p.ParseProgram())
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
func (s *Semantic) declare(name string) {
	peek, ok := s.scopeStack.Peek()
	if !ok {
		return
	}

	*peek = Scope{name: false}
}

// resolve after a variable been resolve we mark it as resolved.
func (s *Semantic) resolve(name string) {
	peek, ok := s.scopeStack.Peek()
	if !ok {
		return
	}

	*peek = Scope{name: true}
}

func (s *Semantic) expectIdentifierDeclare(name string) bool {
	peek, ok := s.scopeStack.Peek()
	if !ok {
		s.NewError("Can't read local variable %s in its own initializer", name)
		return false
	}

	v, ok := (*peek)[name]
	if !ok {
		s.NewError("Identifier %s not exists on current scope", name)
		return false
	}

	if !v {
		s.NewError("Can't read local variable %s in its own initializer", name)
		return false
	}

	return true
}

func (s *Semantic) analyze(node ast.Node) ast.Node {
	switch node := node.(type) {
	case *ast.Program:
		for _, v := range node.Statements {
			s.analyze(v)
		}
	case *ast.ExpressionStatement:
		s.analyze(node.Expression)
	case *ast.Function:
		s.declare(node.Name.Value)
		s.resolve(node.Name.Value)

		s.analyze(node.Body)
	case *ast.BlockStatement:
		s.beginScope()
		for _, stmt := range node.Statements {
			s.analyze(stmt)
		}
		s.endScope()
	case *ast.Identifier:
		s.expectIdentifierDeclare(node.Value)
	case *ast.VarStatement:
		s.declare(node.Name.Value)
		s.analyze(node.Value)
		s.resolve(node.Name.Value)
	}
	return node
}
