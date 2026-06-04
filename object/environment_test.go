package object

import "testing"

func TestEnvironmentGet(t *testing.T) {
	env := NewEnvironment()
	str := &String{Value: "Hello"}
	env.Set("name", str)

	v, ok := env.Get("name")
	if !ok {
		t.Fatalf("Unable to get name from environement")
	}

	stringLiteral, ok := v.(*String)
	if !ok {
		t.Fatalf("Expected to get object.String. Got: %t", v)
	}

	if stringLiteral.Value != "Hello" {
		t.Fatalf("env.Get('name') expected to be %s. Got: %s", "Hello", stringLiteral.Value)
	}
}

func TestNewEnclosedEnvironmentGetOuter(t *testing.T) {
	env := NewEnvironment()
	env.Set("name", &String{Value: "Hello"})
	innerEnv := NewEnclosedEnvironment(env)

	v, ok := innerEnv.Get("name")
	if !ok {
		t.Fatalf("Unable to get name from environement")
	}

	stringLiteral, ok := v.(*String)
	if !ok {
		t.Fatalf("Expected to get object.String. Got: %t", v)
	}

	if stringLiteral.Value != "Hello" {
		t.Fatalf("env.Get('name') expected to be %s. Got: %s", "Hello", stringLiteral.Value)
	}
}

func TestNewEnclosedEnvironmentGetInner(t *testing.T) {
	env := NewEnvironment()
	env.Set("name", &String{Value: "Hello"})
	innerEnv := NewEnclosedEnvironment(env)
	innerEnv.Set("name", &String{Value: "Ninja"})

	v, ok := innerEnv.Get("name")
	if !ok {
		t.Fatalf("Unable to get name from environement")
	}

	stringLiteral, ok := v.(*String)
	if !ok {
		t.Fatalf("Expected to get object.String. Got: %t", v)
	}

	if stringLiteral.Value != "Ninja" {
		t.Fatalf("env.Get('name') expected to be %s. Got: %s", "Ninja", stringLiteral.Value)
	}
}

func TestGetAt(t *testing.T) {
	global := NewEnvironment()
	global.Set("x", &Integer{Value: 1})
	mid := NewEnclosedEnvironment(global)
	mid.Set("y", &Integer{Value: 2})
	inner := NewEnclosedEnvironment(mid)
	inner.Set("z", &Integer{Value: 3})

	if v := inner.GetAt(0, "z"); v.(*Integer).Value != 3 {
		t.Fatalf("GetAt(0) expected 3, got %v", v)
	}
	if v := inner.GetAt(1, "y"); v.(*Integer).Value != 2 {
		t.Fatalf("GetAt(1) expected 2, got %v", v)
	}
	if v := inner.GetAt(2, "x"); v.(*Integer).Value != 1 {
		t.Fatalf("GetAt(2) expected 1, got %v", v)
	}
}

func TestSetAt(t *testing.T) {
	global := NewEnvironment()
	global.Set("x", &Integer{Value: 1})
	mid := NewEnclosedEnvironment(global)
	inner := NewEnclosedEnvironment(mid)

	inner.SetAt(2, "x", &Integer{Value: 99})

	v, ok := global.Get("x")
	if !ok {
		t.Fatalf("expected x in global after SetAt")
	}
	if v.(*Integer).Value != 99 {
		t.Fatalf("SetAt(2) expected global x=99, got %v", v)
	}
}
