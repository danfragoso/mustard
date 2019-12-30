package mustard

import "github.com/tfriedel6/canvas"

type WidgetRenderCallback func(surface *canvas.Canvas)

//SurfaceWidget - Surface Widget
type SurfaceWidget struct {
	widget
	windowSurface  *canvas.Canvas
	renderCallback WidgetRenderCallback
}
