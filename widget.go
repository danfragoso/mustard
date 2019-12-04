package mustard

import (
	"github.com/tfriedel6/canvas"
)

func debugLayout(surface *canvas.Canvas, top, left, width, height int) {
	surface.SetStrokeStyle("#ffa500")
	surface.SetLineWidth(2)
	var dashStyle []float64

	dashStyle = append(dashStyle, 2)
	surface.SetLineDash(dashStyle)
	surface.StrokeRect(float64(left+2), float64(top+2), float64(width-4), float64(height-4))
}

func drawLabelWidget(surface *canvas.Canvas, widget *LabelWidget, top, left, width, height int) {
	surface.SetFillStyle(widget.backgroundColor)
	surface.FillRect(float64(left), float64(top), float64(width), float64(height))

	surface.SetFillStyle(widget.fontColor)
	surface.SetFont("roboto.ttf", widget.fontSize)
	surface.FillText(widget.content, float64(left), float64(top)+widget.fontSize/2+2)

	debugLayout(surface, top, left, width, height)
}
