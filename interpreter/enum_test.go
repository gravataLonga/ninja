package interpreter

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{
			`enum status { case OK:1; case NOK:2;}; status::OK;`,
			1,
		},
		{
			`enum status { case OK: if (true) { 1 } else { 0 }; case NOK:2;}; status::OK;`,
			1,
		},
		{
			`enum status { case OK: "OK"; case NOK:"NOK";}; status::OK;`,
			"OK",
		},
		{
			`enum status { case OK: true; case NOK:false;}; status::OK;`,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEnum[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			testObjectLiteral(t, evaluated, tt.expectedValue)
		})
	}
}
