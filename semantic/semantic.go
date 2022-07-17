package semantic

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
)

// Semantic here is where we are going doing some semantic analysis
// using visitor pattern
type Semantic struct {
	scopeStack      Stack
	globalVariables []string
	localVariables  map[ast.Node]int
	errors          []string
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

func (s *Semantic) Analysis(node *ast.Program) *ast.Program {
	return s.analysis(node).(*ast.Program)
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
func (s *Semantic) declare(name string) {
	if s.scopeStack.IsEmpty() {
		s.globalVariables = append(s.globalVariables, name)
		return
	}

	peek, _ := s.scopeStack.Peek()
	(*peek).Put(name, false)
}

// resolve after a variable been resolve we mark it as resolved.
func (s *Semantic) resolve(name string) {
	peek, ok := s.scopeStack.Peek()
	if !ok {
		return
	}

	(*peek).Put(name, true)
	// s.localVariables[node] =
}

func (s *Semantic) expectIdentifierDeclare(ident *ast.Identifier) bool {
	name := ident.Value
	tok := ident.Token
	peek, ok := s.scopeStack.Peek()
	if !ok {
		// is global environment
		return true
	}

	v, ok := (*peek)[name]
	if !ok {
		// probably is global environment
		return true
	}

	if !v {
		s.NewError("Can't read local variable \"%s\" in its own initializer %s", name, tok)
		return false
	}

	s.localVariables[ident] = s.scopeStack.Size()
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
		s.declare(node.Name.Value)
		s.analysis(node.Value)
		s.resolve(node.Name.Value)
	case *ast.ExpressionStatement:
		s.analysis(node.Expression)
	case *ast.PrefixExpression:
		s.analysis(node.Right)
	case *ast.InfixExpression:
		s.analysis(node.Left)
		s.analysis(node.Right)
	case *ast.FunctionLiteral:
		s.newScope()
		for _, arg := range node.Parameters {
			s.declare(arg.Value)
			s.resolve(arg.Value)
		}
		s.analysis(node.Body)
		s.exitScope()
	case *ast.BlockStatement:
		s.newScope()
		for _, b := range node.Statements {
			s.analysis(b)
		}
		s.exitScope()
	}
	return node
}
