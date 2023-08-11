package interpreter

import (
	"fmt"
	"testing"
)

func TestNewResolver(t *testing.T) {
	input := resolver(t, `function add(x,y) { return x+y; } add(5, add(5, 5));`)

	fmt.Println(input)
}
