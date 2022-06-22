package evaluator

import (
	"fmt"
	"testing"
)

func TestDeleteStatementArray(t *testing.T) {

	evaluated := testEval(`var a = [0, 1]; delete a[0]; a;`, t)
	fmt.Println(evaluated)

}
