package object

import (
	"io"
	"os"
)

var (
	// Arguments is argument that is passed from CLI arguments
	Arguments []string

	// StandardInput where is standard input
	StandardInput io.Reader = os.Stdin

	// StandardOutput where is standard output
	StandardOutput io.Writer = os.Stdout

	// ExitFunction where function responsible for exit
	ExitFunction func(int)
)
