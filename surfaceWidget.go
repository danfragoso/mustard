package mustard

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
)

//CreateSurfaceWidget - Creates and returns a new Surface Widget
func CreateSurfaceWidget(window *Window, frameFunction func(rawSurface *canvas.Canvas)) *SurfaceWidget {
	var widgets []interface{}

	return &SurfaceWidget{
		widget: widget{
			top:  0,
			left: 0,

			width:  0,
			height: 0,

			dirty:   true,
			widgets: widgets,

			padding: 0,

			ref: "button",

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			clickable: false,
			focusable: false,

			backgroundColor: "#fff",
			onClick:         func() {},
		},
		renderCallback: frameFunction,
		windowSurface:  window.surface,
	}
}

//SetWidth - Sets the button width
func (surface *SurfaceWidget) SetWidth(width int) {
	surface.width = width
	surface.fixedWidth = true
}

//SetHeight - Sets the Button height
func (surface *SurfaceWidget) SetHeight(height int) {
	surface.height = height
	surface.fixedHeight = true
}

//SetBackgroundColor - Sets the button background color
func (surface *SurfaceWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		surface.backgroundColor = backgroundColor
		surface.dirty = true
	}
}
