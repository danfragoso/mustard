package mustard

import (
	"github.com/tfriedel6/canvas"
)

//CreateFrame - Creates and returns a new Frame
func CreateFrame(orientation FrameOrientation) *Frame {
	var widgets []interface{}

	return &Frame{
		widget: widget{0, 0, 100, 100, true, widgets, "#fff"},

		orientation: orientation,
	}
}

//AttachWidget - Attaches widgets to a frame el
func (frame *Frame) AttachWidget(widget interface{}) {
	frame.widgets = append(frame.widgets, widget)
}

//SetBackgroundColor - Sets the frame background color
func (frame *Frame) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		frame.backgroundColor = backgroundColor
		frame.dirty = true
	}
}

func drawRootFrame(window *Window) {
	drawFrame(window.surface, window.rootFrame, 0, 0, window.width, window.height)
}

func drawFrame(surface *canvas.Canvas, frame *Frame, top, left, width, height int) {
	surface.SetFillStyle(frame.backgroundColor)
	surface.FillRect(float64(left), float64(top), float64(width), float64(height))

	childrenLen := len(frame.widgets)
	if childrenLen > 0 {
		childrenLayout := calculateChildrenLayout(frame.widgets, top, left, width, height, frame.orientation)

		for i := 0; i < childrenLen; i++ {
			switch frame.widgets[i].(type) {
			case *Frame:
				drawFrame(surface, frame.widgets[i].(*Frame), childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)
			case *LabelWidget:
				drawLabelWidget(surface, frame.widgets[i].(*LabelWidget), childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)
			}
		}
	}

	debugLayout(surface, top, left, width, height)
}

func calculateChildrenLayout(children []interface{}, top, left, width, height int, orientation FrameOrientation) []*widget {
	var childrenLayout []*widget

	childrenLen := len(children)
	for i := 0; i < childrenLen; i++ {
		currentLayout := &widget{}

		if orientation == VerticalFrame {
			currentLayout.top = top
			currentLayout.left = left + (width/childrenLen)*i
			currentLayout.width = width / childrenLen
			currentLayout.height = height
		} else {
			currentLayout.top = top + (height/childrenLen)*i
			currentLayout.left = left
			currentLayout.width = width
			currentLayout.height = height / childrenLen
		}

		childrenLayout = append(childrenLayout, currentLayout)
	}

	return childrenLayout
}
