package mustard

import "github.com/tfriedel6/canvas"

func drawLabelWidget(widget *LabelWidget, surface *canvas.Canvas, top, left int) {
	surface.SetFillStyle(widget.fontColor)
	surface.SetFont("roboto.ttf", widget.fontSize)
	surface.FillText(widget.content, float64(left), widget.fontSize/2+widget.fontSize/10+top)
}

func drawWidgets(widgets []interface{}, surface *canvas.Canvas, top, left int) {
	widgetsLen := len(widgets)

	for i := 0; i < widgetsLen; i++ {
		switch widgets[i].(type) {

		case *LabelWidget:
			widget := widgets[i].(*LabelWidget)
			calculateWidgetsSize(widget, widget.Widgets, i, top, left)

			drawWidgets(widget.widgets, surface, int(widget.top), int(widget.left))
			drawLabelWidget(widget, surface, top, left)
		}
	}
}

func calculateLabelSize(widget *Widget, widgets []interface{}, i, top, left int) {

}
