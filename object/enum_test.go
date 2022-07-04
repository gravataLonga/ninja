package object

import (
	"strings"
	"testing"
)

func TestEnum_Inspect(t *testing.T) {
	expected := "{case STR : Yes; case INT : 1; case FLOAT : 1.000000; case BOOLEAN : false}"
	enum := &Enum{Branches: map[string]Object{}}
	enum.Branches["STR"] = &String{Value: "Yes"}
	enum.Branches["INT"] = &Integer{Value: 1}
	enum.Branches["FLOAT"] = &Float{Value: 1}
	enum.Branches["BOOLEAN"] = FALSE

	if strings.Index(enum.Inspect(), "case STR : Yes") == -1 {
		t.Errorf("enum.inspect() expected to be %s. Got: %s", expected, enum.Inspect())
	}

	if strings.Index(enum.Inspect(), "case INT : 1") == -1 {
		t.Errorf("enum.inspect() expected to be %s. Got: %s", expected, enum.Inspect())
	}

	if strings.Index(enum.Inspect(), "case FLOAT : 1.000000") == -1 {
		t.Errorf("enum.inspect() expected to be %s. Got: %s", expected, enum.Inspect())
	}

	if strings.Index(enum.Inspect(), "case BOOLEAN : false") == -1 {
		t.Errorf("enum.inspect() expected to be %s. Got: %s", expected, enum.Inspect())
	}
}

func TestEnum_Type(t *testing.T) {
	enum := &Enum{Branches: map[string]Object{}}

	if enum.Type() != ENUM_OBJ {
		t.Errorf("enum.type() expected to be %s. Got: %s", ENUM_OBJ, enum.Type())
	}
}
