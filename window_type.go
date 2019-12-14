package mustard

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

//Window - Mustard window
type Window struct {
	visible     bool
	initialized bool

	title string

	width  int
	height int

	top  int
	left int

	rootFrame *Frame
	isDirty   bool

	glw     *glfw.Window
	surface *canvas.Canvas
	backend *goglbackend.GoGLBackend
}
