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

	left.SetBackgroundColor("#ff0000")
	left.SetWidth(30)

	center.SetBackgroundColor("#00ff00")

	right.SetBackgroundColor("#0000ff")
	right.SetWidth(30)

	topBar := CreateFrame(VerticalFrame)
	viewport := CreateFrame(HorizontalFrame)

	viewport.AttachWidget(CreateButtonWidget("Click me!", func() {}))
	viewport.AttachWidget(CreateLabelWidget("memes"))
	viewport.AttachWidget(CreateLabelWidget("memes"))
	viewport.AttachWidget(CreateLabelWidget("memes"))

	topBar.AttachWidget(left)
	topBar.AttachWidget(center)
	topBar.AttachWidget(right)
	topBar.SetHeight(30)

	rootFrame.AttachWidget(topBar)
	rootFrame.AttachWidget(viewport)

	mainWindow.SetRootFrame(rootFrame)
	mainWindow.Show()

	app.Run(func() {

	})

	t.Log(app)
}
