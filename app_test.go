package mustard

import (
	"testing"
)

func TestMain(t *testing.T) {
	app := CreateApp("memes")

	mainWindow := app.CreateWindow("Inicio", 100, 100, 10, 10)

	rootFrame := CreateFrame(HorizontalFrame)

	left := CreateLabelWidget("left")
	center := CreateLabelWidget("center")
	right := CreateLabelWidget("right")

	topBar := CreateFrame(VerticalFrame)
	viewport := CreateFrame(HorizontalFrame)

	topBar.AttachWidget(left)
	topBar.AttachWidget(center)
	topBar.AttachWidget(right)

	rootFrame.AttachWidget(topBar)
	rootFrame.AttachWidget(viewport)

	mainWindow.SetRootFrame(rootFrame)
	mainWindow.Show()

	for index := 0; index < 9999999999; index++ {
	}

	t.Log(app)
}
