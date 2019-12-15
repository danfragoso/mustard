package mustard

func pointerCoordEvent(window *Window, x, y float64) {
	widgets := window.rootFrame.widgets

	focusedWidget := getFocusedWidget(getCoreWidgets(widgets), x, y)
	if focusedWidget != nil {
		window.currentCursor = focusedWidget.cursor
		focusedWidget.focused = true
		window.isDirty = true
	} else {
		window.currentCursor = window.defaultCursor
	}

	window.glw.SetCursor(window.currentCursor)
}

func getFocusedWidget(widgets []*widget, x, y float64) *widget {
	var focusedWidget *widget
	for _, widget := range widgets {
		widget.focused = false

		if len(widget.widgets) > 0 {
			focusedWidget = getFocusedWidget(getCoreWidgets(widget.widgets), x, y)
		} else {
			if calcCursorIntersection(widget, x, y) {
				focusedWidget = widget
			}
		}
	}

	return focusedWidget
}
