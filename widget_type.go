package mustard

import "github.com/go-gl/glfw/v3.2/glfw"

type widget struct {
	top    int
	left   int
	width  int
	height int

	dirty       bool
	fixedWidth  bool
	fixedHeight bool

	widgets []interface{}

	backgroundColor string

	ref    string
	cursor *glfw.Cursor

	padding int

	selectable bool
	focusable  bool
	clickable  bool

	focused  bool
	clicked  bool
	selected bool
}
