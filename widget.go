package mustard

import "github.com/tfriedel6/canvas"

func drawLabelWidget(widget *LabelWidget, surface *canvas.Canvas) {
	surface.SetFillStyle("#000")
	surface.SetFont("roboto.ttf", 20)
	surface.FillText(widget.content, 0, 20/2+2)
}
