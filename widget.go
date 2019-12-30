package mustard

import (
	"github.com/tfriedel6/canvas"
)

func calcCursorIntersection(widget *widget, xF, yF float64) bool {
	x, y := int(xF), int(yF)

	if x > widget.left+widget.padding && x < widget.left+(widget.width-widget.padding) && y > widget.top+widget.padding && y < widget.top+(widget.height-widget.padding) {
		return true
	}

	return false
}

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

	//debugLayout(surface, top, left, width, height)
}

func drawSurfaceWidget(surface *canvas.Canvas, surfaceWidget *SurfaceWidget, top, left, width, height int) {
	surface.FillRect(float64(left), float64(top), float64(width), float64(height))

	surfaceWidget.renderCallback(surfaceWidget.windowSurface)
	debugLayout(surface, top, left, width, height)
}

func drawButtonWidget(surface *canvas.Canvas, widget *ButtonWidget, top, left, width, height int) {
	buttonBackgroundColor := "#f2f2f2"
	buttonBorderColor := "#e0e0e0"
	buttonFontColor := "#444"
	textWidth := surface.MeasureText(widget.content).Width

	if widget.focused {
		buttonBackgroundColor = buttonBorderColor

		//widget.onClick()
	}

	surface.SetFillStyle(buttonBackgroundColor)
	surface.FillRect(float64(left+widget.padding), float64(top+widget.padding), float64(widget.width-widget.padding*2), float64(widget.height-widget.padding*2))

	surface.SetLineWidth(1)
	surface.SetStrokeStyle(buttonBorderColor)
	surface.StrokeRect(float64(left+widget.padding), float64(top+widget.padding), float64(widget.width-widget.padding*2), float64(widget.height-widget.padding*2))

	surface.SetFillStyle(buttonFontColor)
	surface.SetFont("roboto.ttf", 20)
	surface.FillText(widget.content, float64(left+widget.width/2)-textWidth/2, float64(top+widget.height/2+8))

	//debugLayout(surface, top, left, width, height)
}

func drawButtonWidgetB(surface *canvas.Canvas, widget *ButtonWidget, top, left, width, height int) {
	buttonBackgroundColor := "#e7e7e7"
	buttonBorderColor := "#e7e7e7"
	buttonFontColor := "#000"
	buttonPadding := 5.0
	buttonWidth := surface.MeasureText(widget.content).Width + 10

	if widget.focused {
		buttonBackgroundColor = buttonBorderColor
	}

	surface.SetFillStyle(buttonBackgroundColor)
	surface.FillRect(float64(left)-buttonPadding+buttonWidth/2, float64(top+widget.height/2)-buttonPadding, buttonWidth-buttonPadding, float64(widget.height)-buttonPadding)

	surface.SetLineWidth(1)
	surface.SetStrokeStyle(buttonBorderColor)
	surface.StrokeRect(float64(left)-buttonPadding+buttonWidth/2, float64(top+widget.height/2)-buttonPadding, buttonWidth-buttonPadding, float64(widget.height)-buttonPadding)

	surface.SetFillStyle(buttonFontColor)
	surface.SetFont("roboto.ttf", 20)
	surface.FillText(widget.content, float64(left)-buttonPadding+buttonWidth/2+2.5, float64(top+widget.height/2)-buttonPadding+20)

	//debugLayout(surface, top, left, width, height)
}
