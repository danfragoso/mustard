package mustard

import (
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	app := CreateApp("memes")

	mainWindow := app.CreateWindow("Inicio", 100, 100, 10, 10)
	rootFrame := CreateFrame(HorizontalFrame)

	text := CreateLabelWidget("0")
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

	rootFrame.AttachWidget(text)
	rootFrame.AttachWidget(buttons)

	mainWindow.SetRootFrame(rootFrame)
	mainWindow.Show()

	app.Run(func() {

	})

	t.Log(app)
}
