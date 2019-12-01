package mustard

import (
	"github.com/tfriedel6/canvas"
)

func drawLabelWidget(widget *LabelWidget, surface *canvas.Canvas, top, left, width, height int) {
	//surface.SetStrokeStyle("#ffa500")
	///surface.SetLineWidth(3)

	//var dashStyle []float64

	//dashStyle = append(dashStyle, 4)
	//surface.SetLineDash(dashStyle)
	//surface.StrokeRect(float64(left+3), float64(top+3), float64(width-6), float64(height-6))
	surface.SetFillStyle(widget.backgroundColor)
	surface.FillRect(float64(left), float64(top), float64(width), float64(height))

	surface.SetFillStyle(widget.fontColor)
	surface.SetFont("roboto.ttf", widget.fontSize)
	surface.FillText(widget.content, float64(left), float64(top)+widget.fontSize/2+2)
}

func drawWidgets(widgets []interface{}, surface *canvas.Canvas, top, left, width, height int) {
	widgetsLen := len(widgets)

	for i := 0; i < widgetsLen; i++ {
		switch widgets[i].(type) {

		case *LabelWidget:
			widget := widgets[i].(*LabelWidget)
			width, height = calculateWidgetSize(&widget.Widget, width, height)
			drawLabelWidget(widget, surface, top*(i+1), left*(i+1), width, height)

			if widget.Widget.orientation == RowWidgetOrientation {
				drawWidgets(widget.widgets, surface, top*(i+1)+height, left*(i+1), width, height)
			} else {
				drawWidgets(widget.widgets, surface, top*(i+1), left*(i+1)+width, width, height)
			}
		}
	}
}

func calculateWidgetSize(widget *Widget, width, height int) (int, int) {
	childrenWidgetsLen := len(widget.widgets)

	if widget.orientation == RowWidgetOrientation {
		height = height / (childrenWidgetsLen + 1)
	} else {
		width = width / (childrenWidgetsLen + 1)
	}

	return width, height
}
