package mustard

import "github.com/go-gl/glfw/v3.2/glfw"

//CreateApp - Creates and returns a new mustard app
func CreateApp(name string) *App {
	app := App{
		Name: name,
	}

	glfw.Init()
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)
	glfw.WindowHint(glfw.Visible, glfw.False)

	return &app
}

//CreateWindow - Create a new window and adds it to the window app
func (app *App) CreateWindow(windowTitle string, width int, height int, top int, left int) *Window {
	window := newWindow(windowTitle, width, height, top, left)
	app.addWindow(window)
	return window
}

func (app *App) addWindow(window *Window) {
	app.Windows = append(app.Windows, window)
}