package mustard

import "github.com/go-gl/glfw/v3.2/glfw"

//CreateLabelWidget - Creates and returns a new Label Widget
func CreateLabelWidget(content string) *LabelWidget {
	var widgets []interface{}

	return &LabelWidget{
		widget: widget{
			top:  0,
			left: 0,

			width:  0,
			height: 0,

			dirty:   true,
			widgets: widgets,

			ref: "label",

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",
		},
		content: content,

		fontSize:  20,
		fontColor: "#000",
	}
}

//AttachWidget - Attaches a new widget to the window
func (label *LabelWidget) AttachWidget(widget interface{}) {
	label.widgets = append(label.widgets, widget)
}

//SetWidth - Sets the label width
func (label *LabelWidget) SetWidth(width int) {
	label.width = width
	label.fixedWidth = true
}

//SetHeight - Sets the label height
func (label *LabelWidget) SetHeight(height int) {
	label.height = height
	label.fixedHeight = true
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

//SetBackgroundColor - Sets the label background color
func (label *LabelWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		label.backgroundColor = backgroundColor
		label.dirty = true
	}
}
