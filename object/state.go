package object

import "io"

var (
	// Arguments is argument that is passed from CLI arguments
	Arguments []string

	// StandardInput where is standard input
	StandardInput io.Reader

	// StandardOutput where is standard output
	StandardOutput io.Writer

	// ExitFunction where function responsible for exit
	ExitFunction func(int)
)
