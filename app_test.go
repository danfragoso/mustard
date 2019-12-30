package mustard

import (
	"strconv"
	"testing"

	"github.com/tfriedel6/canvas"
)

func TestMain(t *testing.T) {
	app := CreateApp("memes")

	mainWindow := app.CreateWindow("Inicio", 100, 100, 10, 10)
	rootFrame := CreateFrame(HorizontalFrame)

	text := CreateLabelWidget("0")
	rawSurface := CreateSurfaceWidget(mainWindow, drawBrowserFrame)
	buttons := CreateFrame(VerticalFrame)

	buttons.SetHeight(100)

	plusOne := CreateButtonWidget("+1", func() {
		val, _ := strconv.Atoi(text.GetContent())
		val = val + 1
		text.SetContent(strconv.Itoa(val))
	})

	minusOne := CreateButtonWidget("-1", func() {
		val, _ := strconv.Atoi(text.GetContent())
		val = val - 1
		text.SetContent(strconv.Itoa(val))
	})

	buttons.AttachWidget(plusOne)
	buttons.AttachWidget(minusOne)

	rootFrame.AttachWidget(rawSurface)
	rootFrame.AttachWidget(text)
	rootFrame.AttachWidget(buttons)

	mainWindow.SetRootFrame(rootFrame)
	mainWindow.Show()

	app.Run(func() {
		t.Log("memes")
	})

	t.Log(app)
}

func drawBrowserFrame(surface *canvas.Canvas) {
	surface.SetFillStyle("#ff0000")
	surface.FillRect(10, 10, 20, 20)
}
