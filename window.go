package mustard

//NewWindow - Creates and returns a new mustard window
func NewWindow(windowTitle string, width int, height int, top int, left int) *Window {
	window := Window{}

	window.Visible = false
	window.Initialized = false

	window.Title = windowTitle

	window.Width = width
	window.Height = height

	window.Top = top
	window.Left = left

	return &window
}
