package mustard

//CreateButtonWidget - Creates and returns a new Button Widget
func CreateButtonWidget(content string) *ButtonWidget {
	var widgets []interface{}

	return &ButtonWidget{
		widget: widget{
			top:  0,
			left: 0,

			width:  0,
			height: 30,

			dirty:   true,
			widgets: widgets,

			clickable: true,
			focusable: true,

			backgroundColor: "#fff",
		},
		content: content,
	}
}

//AttachWidget - Attaches a new widget to the window
func (button *ButtonWidget) AttachWidget(widget interface{}) {
	button.widgets = append(button.widgets, widget)
}

//SetWidth - Sets the button width
func (button *ButtonWidget) SetWidth(width int) {
	button.width = width
	button.fixedWidth = true
}

//SetHeight - Sets the Button height
func (button *ButtonWidget) SetHeight(height int) {
	button.height = height
	button.fixedHeight = true
}

//SetBackgroundColor - Sets the button background color
func (button *ButtonWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		button.backgroundColor = backgroundColor
		button.dirty = true
	}
}
