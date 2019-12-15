package mustard

func pointerCoordEvent(window *Window, x, y float64) {
	widgets := window.rootFrame.widgets

	getFocusedWidget(window, widgets, x, y)
}

func getFocusedWidget(window *Window, widgets []interface{}, x, y float64) {
	for i := 0; i < len(widgets); i++ {
		switch widgets[i].(type) {
		case *Frame:
			widget := widgets[i].(*Frame)
			if widget.focusable {
				widget.cursorIntersection(window, x, y)
			}
		case *LabelWidget:
			widget := widgets[i].(*LabelWidget)
			if widget.focusable {
				widget.cursorIntersection(window, x, y)
			}
		case *ButtonWidget:
			widget := widgets[i].(*ButtonWidget)
			if widget.focusable {
				widget.cursorIntersection(window, x, y)
			}
		}
	}
}
