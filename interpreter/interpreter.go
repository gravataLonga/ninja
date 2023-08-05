package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"github.com/gravataLonga/ninja/stdlib"
	"io"
	"os"
	"strings"
)

type Interpreter struct {
	env       *object.Environment
	globals   *object.Environment
	innerLoop int
	locals    map[ast.Expression]int
	output    io.Writer
}

func New(w io.Writer, env *object.Environment) *Interpreter {
	env.Set("len", object.NewBuiltin(stdlib.Len))
	env.Set("first", object.NewBuiltin(stdlib.First))
	env.Set("puts", object.NewBuiltin(stdlib.Puts))
	env.Set("last", object.NewBuiltin(stdlib.Last))
	env.Set("rest", object.NewBuiltin(stdlib.Rest))
	env.Set("push", object.NewBuiltin(stdlib.Push))
	env.Set("time", object.NewBuiltin(stdlib.Time))
	env.Set("rand", object.NewBuiltin(stdlib.Rand))
	env.Set("args", object.NewBuiltin(stdlib.Args))
	env.Set("plugin", object.NewBuiltin(stdlib.Plugin))

	return &Interpreter{
		env:    env,
		output: w,
		locals: make(map[ast.Expression]int),
	}
}

func (i *Interpreter) EnterLoop() {
	i.innerLoop++
}

func (i *Interpreter) ExitLoop() {
	i.innerLoop--
}

func (i *Interpreter) InLoop() bool {
	return i.innerLoop > 0
}

func (i *Interpreter) evaluate(node ast.Expression) object.Object {
	return node.Accept(i)
}

func (i *Interpreter) Interpreter(node ast.Node) object.Object {
	return i.execute(node)
}

func (i *Interpreter) execute(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		result := node.Accept(i)
		return result
	case *ast.BlockStatement:
		result := node.Accept(i)
		return result
	case *ast.ExpressionStatement:
		result := node.Accept(i)
		return result
	case *ast.ReturnStatement:
		result := node.Accept(i)
		return result
	case *ast.BreakStatement:
		result := node.Accept(i)
		return result
	case *ast.VarStatement:
		result := node.Accept(i)
		return result
	case *ast.AssignStatement:
		result := node.Accept(i)
		return result
	}
	return nil
}

// Statements

func (i *Interpreter) VisitProgram(v *ast.Program) (result object.Object) {
	for _, stmt := range v.Statements {
		result = stmt.Accept(i)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Break:
			return object.NewErrorFormat("'break' not in the 'loop' context")
		case *object.Error:
			return result
		}
	}
	return
}

func (i *Interpreter) VisitBlock(v *ast.BlockStatement) (result object.Object) {
	for _, stmt := range v.Statements {
		result = i.execute(stmt)
		if result == nil {
			continue
		}

		// @todo test it better
		if i.InLoop() {
			if object.IsReturn(result) {
				return
			}

			if result.Type() == object.BREAK_VALUE_OBJ {
				return
			}
		}

		if object.IsReturn(result) {
			return result.(*object.ReturnValue).Value
		}

		if object.IsError(result) {
			return
		}
	}
	return
}

func (i *Interpreter) VisitBreak(v *ast.BreakStatement) (result object.Object) {
	return &object.Break{Value: object.NULL}
}

func (i *Interpreter) VisitDelete(v *ast.DeleteStatement) (result object.Object) {
	ident, ok := v.Left.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left must be a identifier. Got: %T", v.Left)
	}

	value, ok := i.env.Get(ident.Value)
	if !ok {
		return object.NewErrorFormat("DeleteStatement.left %s identifier not found.", ident.Value)
	}

	index := i.evaluate(v.Index)

	switch value.(type) {
	case *object.Array:
		arr, _ := value.(*object.Array)
		if !object.IsInteger(index) {
			return object.NewErrorFormat("DeleteStatement.index must be a Integer. Got: %T", index)
		}
		index, _ := index.(*object.Integer)
		arr.Elements = removeIndexFromArray(arr.Elements, index.Value)
		i.env.Set(ident.Value, arr)
	case *object.Hash:
		hash, _ := value.(*object.Hash)
		hashable, ok := index.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("DeleteStatement.index must be a Hashable. Got: %T", index)
		}
		delete(hash.Pairs, hashable.HashKey())

		i.env.Set(ident.Value, hash)
	default:
		return object.NewErrorFormat("DeleteStatement.left only work with array or hash object. Got: %T", value)
	}

	return nil
}

// removeIndexFromArray is slow operations, we need better way?
func removeIndexFromArray(slice []object.Object, s int64) []object.Object {

	copy(slice[s:], slice[s+1:])      // Shift a[i+1:] left one index.
	slice[len(slice)-1] = object.NULL // Erase last element (write zero value).
	slice = slice[:len(slice)-1]      // Truncate slice.

	return slice
}

func (i *Interpreter) VisitEnum(v *ast.EnumStatement) (result object.Object) {
	enum := &object.Enum{Branches: map[string]object.Object{}}
	for o, v := range v.Branches {
		enum.Branches[o] = i.evaluate(v)
	}

	ident, ok := v.Identifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected identifier. got: %s", v.Identifier)
	}

	i.env.Set(ident.Value, enum)
	return enum
}

func (i *Interpreter) VisitExprStmt(v *ast.ExpressionStatement) (result object.Object) {
	return i.evaluate(v.Expression)
}

func (i *Interpreter) VisitReturn(v *ast.ReturnStatement) (result object.Object) {
	if v.ReturnValue == nil {
		return &object.ReturnValue{Value: object.NULL}
	}

	result = i.evaluate(v.ReturnValue)
	if object.IsError(result) {
		return
	}

	return &object.ReturnValue{Value: result}
}

func (i *Interpreter) VisitVarStmt(v *ast.VarStatement) (result object.Object) {
	i.env.Set(v.Name.Value, i.evaluate(v.Value))
	return nil
}

func (i *Interpreter) VisitAssignStmt(v *ast.AssignStatement) (result object.Object) {
	ident, ok := v.Left.(*ast.Identifier)
	if ok {
		left := ident.Value
		i.env.Set(left, i.evaluate(v.Right))
		return nil
	}

	expr, ok := v.Left.(*ast.ExpressionStatement)
	if !ok {
		return nil
	}

	idx, ok := expr.Expression.(*ast.IndexExpression)
	if !ok {
		return nil
	}

	ident, ok = idx.Left.(*ast.Identifier)
	if !ok {
		return nil
	}

	left := ident.Value

	obj, ok := i.env.Get(left)
	if !ok {
		return nil
	}

	if object.IsArray(obj) {
		arr, _ := obj.(*object.Array)
		index := i.evaluate(idx.Index)
		indexIntegerObject, ok := index.(*object.Integer)
		if !ok {
			return nil
		}

		indexInteger := int(indexIntegerObject.Value)
		lenElements := len(arr.Elements)

		if indexInteger <= -1 {
			return object.NewErrorFormat("index out of range, got %d not positive index", indexInteger)
		}

		if lenElements < indexInteger {
			return object.NewErrorFormat("index out of range, got %d but array has only %d elements", indexInteger, lenElements)
		}

		if indexInteger > lenElements-1 {
			lenElements = lenElements + 1
		}

		elements := make([]object.Object, lenElements)
		copy(elements, arr.Elements)
		elements[indexInteger] = i.evaluate(v.Right)
		arr.Elements = elements
		i.env.Set(left, arr)
	}

	if object.IsHash(obj) {
		hashObject, _ := obj.(*object.Hash)

		objIndex := i.evaluate(idx.Index)
		h, ok := objIndex.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("expected index to be hashable")
		}
		hashObject.Pairs[h.HashKey()] = object.HashPair{Key: objIndex, Value: i.evaluate(v.Right)}
	}

	return nil
}

// Expresions

func (i *Interpreter) VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object) {
	elements := i.evaluateExpressions(v.Elements)
	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

func (i *Interpreter) VisitBooleanExpr(v *ast.Boolean) (result object.Object) {
	if v.Value {
		return object.TRUE
	}
	return object.FALSE
}

// VisitCallExpr
// @todo arguments must be it's on AST structure.
func (i *Interpreter) VisitCallExpr(v *ast.CallExpression) (result object.Object) {
	obj := i.evaluate(v.Function)

	if object.IsError(obj) {
		return obj
	}

	switch obj.Type() {
	case object.FUNCTION_OBJ:
		return i.applyFunction(obj, v)
	case object.BUILTIN_OBJ:
		var args []object.Object

		for _, e := range v.Arguments {
			evaluated := i.evaluate(e)
			if object.IsError(evaluated) {
				// return []object.Object{evaluated}
			}
			args = append(args, evaluated)
		}

		return obj.(*object.Builtin).Fn(args...)
	default:
		return object.NewErrorFormat("Not implement yet VisitCallExpr")
	}
}

func (i *Interpreter) applyFunction(obj object.Object, v *ast.CallExpression) (result object.Object) {
	fn, _ := obj.(*object.FunctionLiteral)
	parameters := fn.Parameters.([]ast.Expression)

	mParameter := len(v.Arguments)
	err := i.validateArguments(v, parameters)

	if object.IsError(err) {
		return err
	}

	envLocal := object.NewEnclosedEnvironment(i.env)

	for index, parameter := range parameters {
		ident, ok := parameter.(*ast.Identifier)
		if ok {
			envLocal.Set(ident.String(), i.evaluate(v.Arguments[index]))
			continue
		}

		infix, ok := parameter.(*ast.InfixExpression)
		ident, _ = infix.Left.(*ast.Identifier)

		var value object.Object
		if mParameter > index {
			argument := v.Arguments[index]
			value = i.evaluate(argument)
		} else {
			value = i.evaluate(infix.Right)
		}

		envLocal.Set(ident.String(), value)
	}

	env := i.env
	i.env = envLocal
	result = i.VisitBlock(fn.Body.(*ast.BlockStatement))
	i.env = env

	if object.IsReturn(result) {
		return result.(*object.ReturnValue).Value
	}
	return
}

func (i *Interpreter) validateArguments(v *ast.CallExpression, parameters []ast.Expression) object.Object {
	// cases:
	// 1. fn () {}(1);
	// 2. fn (x) {}();
	// 3. fn (x) {}(1, 2);
	// 4. fn (x, y = 1) {}();
	// 5. fn (x, y = 1) {}(1,2,3);

	/*
		A parameter is a variable in a function definition. It is a placeholder and hence does not have a concrete value.
		An argument is a value passed during function invocation.
	*/

	totalParameters := len(parameters)
	totalParametersDefault := 0
	totalParametersRequireBeforeDefault := 0
	totalArguments := len(v.Arguments)

	for _, p := range parameters {
		if _, ok := p.(*ast.InfixExpression); ok {
			totalParametersDefault++
		}
		if _, ok := p.(*ast.Identifier); ok {
			totalParametersRequireBeforeDefault++
		}
	}

	if totalParametersDefault > 0 {
		if totalParametersDefault != totalParameters-totalParametersRequireBeforeDefault {
			return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParameters, totalArguments, v.Token)
		}
		if totalParametersDefault+totalParametersRequireBeforeDefault != totalParameters {
			return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParametersDefault+totalParametersRequireBeforeDefault, totalArguments, v.Token)
		}

		return nil
	}

	if totalArguments != totalParameters {
		return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParameters, totalArguments, v.Token)
	}

	return nil
}

func (i *Interpreter) VisitDotExpr(v *ast.Dot) (result object.Object) {
	obj := i.evaluate(v.Object)
	call, ok := v.Right.(*ast.CallExpression)
	if !ok {
		return object.NewErrorFormat("we expect to be a call on right of dot operation. Got: %t", v.Right)
	}

	objCallable, ok := obj.(object.CallableMethod)
	if !ok {
		return object.NewErrorFormat("object must implement callable.")
	}

	method, ok := call.Function.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("method name isn't a identifier")
	}

	args := i.evaluateExpressions(call.Arguments)

	return objCallable.Call(method.Value, args...)

	/*
		function := Eval(node.Function, env)
		if object.IsError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

		return object.NewErrorFormat("Not implement yet VisitDotExpr")
	*/
	return object.NewErrorFormat("Not implement yet VisitDotExpr")
}

func (i *Interpreter) VisitFloatExpr(v *ast.FloatLiteral) (result object.Object) {
	return &object.Float{Value: v.Value}
}

func (i *Interpreter) VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object) {
	fn := &object.FunctionLiteral{Parameters: v.Parameters, Body: v.Body}
	if v.Name != nil {
		i.env.Set(v.Name.Value, fn)
	}
	return fn
}

func (i *Interpreter) VisitHashExpr(v *ast.HashLiteral) (result object.Object) {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range v.Pairs {
		key := i.evaluate(keyNode)
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("unusable as hash key: %s", key.Type())
		}

		value := i.evaluate(valueNode)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func (i *Interpreter) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	value, ok := i.env.Get(v.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: %s %s", v.Value, v.Token)
		// return object.NULL
	}
	return value
}

func (i *Interpreter) VisitIfExpr(v *ast.IfExpression) (result object.Object) {
	// Probably problem is here:
	condition := i.evaluate(v.Condition)
	if object.IsTruthy(condition) {
		return i.execute(v.Consequence)
	}
	if v.Alternative != nil {
		return i.execute(v.Alternative)
	}
	return object.NULL
}

func (i *Interpreter) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object) {
	access, ok := v.AccessIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected access identifier. got: %s", v.AccessIdentifier)
	}

	property, ok := v.PropertyIdentifier.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("expected property identifier. got: %s", v.PropertyIdentifier)
	}

	obj, ok := i.env.Get(access.Value)
	if !ok {
		return object.NewErrorFormat("identifier not found: " + access.Value)
	}

	enum, ok := obj.(*object.Enum)
	if !ok {
		return object.NewErrorFormat("identifier must be accessible with :: got: %s", v)
	}

	brancheValue, ok := enum.Branches[property.Value]
	if !ok {
		return object.NewErrorFormat("identifier %s don't exists on enum object", property.Value)
	}

	return brancheValue
}

func (i *Interpreter) VisitImportExpr(v *ast.Import) (result object.Object) {

	resultFilename := i.evaluate(v.Filename)

	filename, ok := resultFilename.(*object.String)
	if !ok {
		return object.NULL
	}

	readFile, err := os.Open(filename.Value)

	if err != nil {
		return object.NewErrorFormat("IO Error: error reading file '%s': %s %s", filename.Value, err, v.Token)
	}

	l := lexer.New(readFile)
	p := parser.New(l)
	programs := p.ParseProgram()

	if len(p.Errors()) > 0 {
		strErros := []string{}
		for _, e := range p.Errors() {
			strErros = append(strErros, e)
		}
		return object.NewErrorFormat("%s: %s", filename.Value, strings.Join(strErros, "\n"))
	}

	result = i.execute(programs)

	if result == nil {
		return nil
	}

	errorStr, ok := result.(*object.Error)
	if ok {
		return object.NewErrorFormat("%s: %s", filename.Value, errorStr.Message)
	}

	// Only return if last item of imported file have "return"
	if len(programs.Statements) > 0 {
		stmts := programs.Statements
		stmt := stmts[len(stmts)-1]

		_, ok = stmt.(*ast.ReturnStatement)
		if ok {
			return result
		}
	}

	return object.NULL
}

func (i *Interpreter) VisitIndexExpr(v *ast.IndexExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	index := i.evaluate(v.Index)
	return indexExpression(left, index)
}

func (i *Interpreter) VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object) {
	return &object.Integer{Value: v.Value}
}

func (i *Interpreter) VisitPostfixExpr(v *ast.PostfixExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	if object.IsError(left) {
		return left
	}

	result = postfixExpression(v, left)
	astIdent, ok := v.Left.(*ast.Identifier)
	if !ok {
		return
	}

	ident := astIdent.Token
	i.env.Set(ident.Literal, result)
	return left
}

func (i *Interpreter) VisitPrefixExpr(v *ast.PrefixExpression) (result object.Object) {
	right := i.evaluate(v.Right)
	if object.IsError(right) {
		return right
	}

	result = prefixExpression(v, right)
	astIdent, ok := v.Right.(*ast.Identifier)
	if !ok {
		return
	}

	ident := astIdent.Token
	i.env.Set(ident.Literal, result)
	return
}

func (i *Interpreter) VisitStringExpr(v *ast.StringLiteral) (result object.Object) {
	return &object.String{Value: v.Value}
}

func (i *Interpreter) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object) {
	condition := i.evaluate(v.Condition)
	if object.IsTruthy(condition) {
		return i.evaluate(v.Consequence)
	}

	return i.evaluate(v.Alternative)
}

func (i *Interpreter) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	if object.IsTruthy(left) {
		return left
	}
	return i.evaluate(v.Right)
}

func (i *Interpreter) VisitFor(v *ast.ForStatement) (result object.Object) {
	if v.InitialCondition != nil {
		i.execute(v.InitialCondition)
	}

	// @todo test this better
	i.EnterLoop()
	condition := i.interpreterConditionForLoop(v.Condition)
	for object.IsTruthy(condition) {

		result = i.execute(v.Body)
		if result != nil {
			if object.IsReturn(result) {
				i.ExitLoop()
				return
			}

			if result.Type() == object.BREAK_VALUE_OBJ {
				i.ExitLoop()
				return nil
			}

			if object.IsError(result) {
				i.ExitLoop()
				return
			}
		}
		if v.Iteration != nil {
			i.execute(v.Iteration)
		}
		condition = i.interpreterConditionForLoop(v.Condition)
	}
	// @todo test this better
	i.ExitLoop()

	return result
}

// interpreterConditionForLoop @todo check better way.
func (i *Interpreter) interpreterConditionForLoop(v ast.Expression) object.Object {
	if v == nil {
		return object.TRUE
	}
	return i.evaluate(v)
}

func (i *Interpreter) VisitInfix(v *ast.InfixExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	right := i.evaluate(v.Right)

	if object.IsError(left) {
		return left
	}

	if object.IsError(right) {
		return right
	}

	return infixExpression(v, v.Operator, left, right)
}

func (i *Interpreter) evaluateExpressions(exprs []ast.Expression) []object.Object {
	var result []object.Object

	for _, e := range exprs {
		evaluated := i.evaluate(e)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}
