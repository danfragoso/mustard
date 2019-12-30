package mustard

import (
	"github.com/tfriedel6/canvas"
)

//CreateFrame - Creates and returns a new Frame
func CreateFrame(orientation FrameOrientation) *Frame {
	var widgets []interface{}

	return &Frame{
		widget: widget{
			top:  0,
			left: 0,

			width:  100,
			height: 100,

			ref: "frame",

			dirty:   true,
			widgets: widgets,

			backgroundColor: "#fff"},

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

//SetWidth - Sets the frame width
func (frame *Frame) SetWidth(width int) {
	frame.width = width
	frame.fixedWidth = true
}

//SetHeight - Sets the frame height
func (frame *Frame) SetHeight(height int) {
	frame.height = height
	frame.fixedHeight = true
}

func drawRootFrame(window *Window) {
	drawFrame(window.surface, window.rootFrame, 0, 0, window.width, window.height)
}

func drawFrame(surface *canvas.Canvas, frame *Frame, top, left, width, height int) {
	surface.SetFillStyle(frame.backgroundColor)
	surface.FillRect(float64(left), float64(top), float64(width), float64(height))

	childrenLen := len(frame.widgets)
	if childrenLen > 0 {
		childrenWidgets := getCoreWidgets(frame.widgets)
		childrenLayout := calculateChildrenWidgetsLayout(childrenWidgets, top, left, width, height, frame.orientation)

		for i := 0; i < childrenLen; i++ {
			switch frame.widgets[i].(type) {
			case *Frame:
				frame := frame.widgets[i].(*Frame)
				frame.top = childrenLayout[i].top
				frame.left = childrenLayout[i].left
				frame.width = childrenLayout[i].width
				frame.height = childrenLayout[i].height

				drawFrame(surface, frame, childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)
			case *SurfaceWidget:
				surfaceWidget := frame.widgets[i].(*SurfaceWidget)
				surfaceWidget.top = childrenLayout[i].top
				surfaceWidget.left = childrenLayout[i].left
				surfaceWidget.width = childrenLayout[i].width
				surfaceWidget.height = childrenLayout[i].height

				drawSurfaceWidget(surface, surfaceWidget, childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)

			case *LabelWidget:
				label := frame.widgets[i].(*LabelWidget)
				label.top = childrenLayout[i].top
				label.left = childrenLayout[i].left
				label.width = childrenLayout[i].width
				label.height = childrenLayout[i].height

				drawLabelWidget(surface, label, childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)
			case *ButtonWidget:
				button := frame.widgets[i].(*ButtonWidget)
				button.top = childrenLayout[i].top
				button.left = childrenLayout[i].left
				button.width = childrenLayout[i].width
				button.height = childrenLayout[i].height

				drawButtonWidget(surface, button, childrenLayout[i].top, childrenLayout[i].left, childrenLayout[i].width, childrenLayout[i].height)
			}
		}
	}

	//debugLayout(surface, top, left, width, height)
}

func getCoreWidgets(widgets []interface{}) []*widget {
	var coreWidgets []*widget

	for i := 0; i < len(widgets); i++ {
		switch widgets[i].(type) {
		case *SurfaceWidget:
			widget := widgets[i].(*SurfaceWidget)
			coreWidgets = append(coreWidgets, &widget.widget)
		case *Frame:
			widget := widgets[i].(*Frame)
			coreWidgets = append(coreWidgets, &widget.widget)
		case *LabelWidget:
			widget := widgets[i].(*LabelWidget)
			coreWidgets = append(coreWidgets, &widget.widget)
		case *ButtonWidget:
			widget := widgets[i].(*ButtonWidget)
			coreWidgets = append(coreWidgets, &widget.widget)
		}
	}
	return coreWidgets
}

func calculateChildrenWidgetsLayout(children []*widget, top, left, width, height int, orientation FrameOrientation) []*widget {
	var childrenLayout []*widget

	childrenLen := len(children)
	for i := 0; i < childrenLen; i++ {
		currentLayout := &widget{}

		if orientation == VerticalFrame {
			fixedWidthElemens, volatileWidthElements := getFixedWidthElements(children)
			remainingWidth := calculateFlexibleWidth(width, fixedWidthElemens)

			if i > 0 {
				currentLayout.left = childrenLayout[i-1].left + childrenLayout[i-1].width
			} else {
				currentLayout.left = left
			}

			if children[i].fixedWidth {
				currentLayout.width = children[i].width
			} else {
				currentLayout.width = remainingWidth / len(volatileWidthElements)
			}

			currentLayout.top = top
			currentLayout.height = height
		} else {
			fixedHeightElements, volatileHeightElements := getFixedHeightElements(children)
			remainingHeight := calculateFlexibleHeight(height, fixedHeightElements)

			if i > 0 {
				currentLayout.top = childrenLayout[i-1].top + childrenLayout[i-1].height
			} else {
				currentLayout.top = top
			}

			if children[i].fixedHeight {
				currentLayout.height = children[i].height
			} else {
				currentLayout.height = remainingHeight / len(volatileHeightElements)
			}

			currentLayout.left = left
			currentLayout.width = width
		}

		childrenLayout = append(childrenLayout, currentLayout)
	}

	return childrenLayout
}

func getFixedWidthElements(elements []*widget) ([]*widget, []*widget) {
	var fixedWidth []*widget
	var volatileWidth []*widget

	for _, element := range elements {
		if element.fixedWidth {
			fixedWidth = append(fixedWidth, element)
		} else {
			volatileWidth = append(volatileWidth, element)
		}
	}
	return fixedWidth, volatileWidth
}

func getFixedHeightElements(elements []*widget) ([]*widget, []*widget) {
	var fixedHeight []*widget
	var volatileHeight []*widget

	for _, element := range elements {
		if element.fixedHeight {
			fixedHeight = append(fixedHeight, element)
		} else {
			volatileHeight = append(volatileHeight, element)
		}
	}
	return fixedHeight, volatileHeight
}

func calculateFlexibleWidth(avaiableWidth int, elements []*widget) int {
	for _, el := range elements {
		avaiableWidth = avaiableWidth - el.width
	}

	return avaiableWidth
}

func calculateFlexibleHeight(avaiableHeight int, elements []*widget) int {
	for _, el := range elements {
		avaiableHeight = avaiableHeight - el.height
	}

	return avaiableHeight
}
