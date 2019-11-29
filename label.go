package mustard

//CreateLabelWidget - Creates and returns a new Label Widget
func CreateLabelWidget(content string) *LabelWidget {
	return &LabelWidget{
		Widget:  Widget{0, 0, 100, 100},
		content: content,
	}
}
