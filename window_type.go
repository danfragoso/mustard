package mustard

import (
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
	"github.com/veandco/go-sdl2/sdl"
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
	close     bool

	sdlw    *sdl.Window
	sdlwID  uint32
	surface *canvas.Canvas

	backend   *goglbackend.GoGLBackend
	glContext sdl.GLContext

	frames int64
}
