package object

import "testing"

func TestEnvironment_Get(t *testing.T) {
	env := NewEnvironment()
	env.Set("name", &String{Value: "Hello"})

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

func TestNewEnclosedEnvironment_GetOuter(t *testing.T) {
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

func TestNewEnclosedEnvironment_GetInner(t *testing.T) {
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
