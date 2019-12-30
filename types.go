package mustard

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

type WidgetRenderCallback func(surface *canvas.Canvas)
type WidgetClickCallback func()

type App struct {
	Name string

	windows []*Window
}

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

	defaultCursor *glfw.Cursor
	currentCursor *glfw.Cursor

	currentFocusedWidget *widget

	glw     *glfw.Window
	surface *canvas.Canvas
	backend *goglbackend.GoGLBackend
}

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

//SurfaceWidget - Surface Widget
type SurfaceWidget struct {
	widget
	windowSurface  *canvas.Canvas
	renderCallback WidgetRenderCallback
}

type LabelWidget struct {
	widget
	content string

	fontSize  float64
	fontColor string
}

type ButtonWidget struct {
	widget

	content string
}

type FrameOrientation int

const (
	//VerticalFrame - Vertical frame orientation
	VerticalFrame FrameOrientation = iota

	//HorizontalFrame - Horizontal frame orientation
	HorizontalFrame
)

//Frame - Layout frame type
type Frame struct {
	widget

	orientation FrameOrientation
}
