package mustard

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

//Window - Mustard window
type Window struct {
	Visible     bool
	Initialized bool

	Title string

	Width  int
	Height int

	Top  int
	Left int

	Widgets []interface{}
	isDirty bool

	glw     *glfw.Window
	surface *canvas.Canvas
	backend *goglbackend.GoGLBackend
}
