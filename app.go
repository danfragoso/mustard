package mustard

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

//CreateApp - Creates and returns a new mustard app
func CreateApp(name string) *App {
	runtime.LockOSThread()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Printf("Error initializing SDL: %v", err)
	}

	app := App{
		Name: name,
	}

	sdl.GLSetAttribute(sdl.GL_RED_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 0)
	sdl.GLSetAttribute(sdl.GL_STENCIL_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GLSetAttribute(sdl.GL_MULTISAMPLEBUFFERS, 1)
	sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4)

	return &app
}

//CreateWindow - Create a new window and adds it to the window app
func (app *App) CreateWindow(windowTitle string, width int, height int, top int, left int) *Window {
	window := newWindow(windowTitle, width, height, top, left)
	app.addWindow(window)
	return window
}

//Run - Run the app
func (app *App) Run(callback func()) {
	for _, window := range app.windows {
		if window.visible {

			if !window.close {
				window.processFrame()
			}
		}
	}

	callback()
	app.Run(callback)
}

func (app *App) addWindow(window *Window) {
	app.windows = append(app.windows, window)
}
