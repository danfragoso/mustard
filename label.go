package mustard

//CreateLabelWidget - Creates and returns a new Label Widget
func CreateLabelWidget(content string) *LabelWidget {
	var widgets []interface{}

	return &LabelWidget{
		Widget:  Widget{0, 0, 100, 100, true, widgets, RowWidgetOrientation},
		content: content,

		fontSize:  20,
		fontColor: "#000",
	}
}

//AttachWidget - Attaches a new widget to the window
func (label *LabelWidget) AttachWidget(widget interface{}) {
	label.widgets = append(label.widgets, widget)
}

//SetFontSize - Sets the label font size
func (label *LabelWidget) SetFontSize(fontSize float64) {
	label.fontSize = fontSize
	label.dirty = true
}

//SetFontColor - Sets the label font color
func (label *LabelWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		label.fontColor = fontColor
		label.dirty = true
	}
}
