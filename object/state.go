package object

import "io"

var (
	Arguments      []string
	StandardInput  io.Reader
	StandardOutput io.Writer
	ExitFunction   func(int)
)
