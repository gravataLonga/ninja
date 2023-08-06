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
