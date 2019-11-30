package mustard

import (
	"testing"
)

func TestMain(t *testing.T) {
	app := CreateApp("memes")

	mainWindow := app.CreateWindow("Inicio", 100, 100, 10, 10)

	A := CreateLabelWidget("A")
	B := CreateLabelWidget("B")
	C := CreateLabelWidget("C")
	D := CreateLabelWidget("D")
	E := CreateLabelWidget("E")

	A.AttachWidget(B)
	A.AttachWidget(C)
	C.AttachWidget(D)
	C.SetLayoutOrientation(ColumnWidgetOrientation)
	D.AttachWidget(E)
	D.SetLayoutOrientation(ColumnWidgetOrientation)
	E.AttachWidget(CreateLabelWidget("F"))
	E.SetLayoutOrientation(ColumnWidgetOrientation)
	mainWindow.AttachWidget(A)
	mainWindow.Show()

	for index := 0; index < 9999999999; index++ {
	}

	t.Log(app)
}
