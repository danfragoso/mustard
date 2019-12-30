package mustard

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

//ButtonClickCallback - The click event callback
type WidgetClickCallback func()

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

	onClick WidgetClickCallback
}
