package mustard

import "github.com/go-gl/glfw/v3.2/glfw"

//Window - Mustard window
type Window struct {
	Visible     bool
	Initialized bool

	Title string

	Width  int
	Height int

	Top  int
	Left int

	glw *glfw.Window
}
