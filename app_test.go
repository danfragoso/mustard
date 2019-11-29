package mustard

import (
	"testing"
)

func TestMain(t *testing.T) {
	app := CreateApp("memes")

	mainWindow := app.CreateWindow("Inicio", 100, 100, 10, 10)

	label := CreateLabelWidget("memes")

	label.SetFontColor("#0f0")
	label.SetFontSize(20)
	label.AttachWidget(CreateLabelWidget("oof"))

	mainWindow.AttachWidget(label)
	mainWindow.Show()

	for index := 0; index < 9999999999; index++ {
	}

	t.Log(app)
}
