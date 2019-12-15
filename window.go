package mustard

import (
	"log"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

func newWindow(windowTitle string, width int, height int, top int, left int) *Window {
	window := Window{}

	window.visible = false
	window.initialized = false

	window.title = windowTitle

	window.width = width
	window.height = height

	window.top = top
	window.left = left

	window.defaultCursor = glfw.CreateStandardCursor(glfw.ArrowCursor)
	window.currentCursor = window.defaultCursor

	glw, err := glfw.CreateWindow(window.width, window.height, window.title, nil, nil)
	if err != nil {
		log.Fatalf("Error creating window: %v", err)
	}

	glw.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}

	window.glw = glw
	window.backend = backend
	window.surface = canvas.New(backend)

	window.glw.SetSizeCallback(func(w *glfw.Window, width, height int) {
		window.width = width
		window.height = height

		window.backend.SetBounds(0, 0, window.width, window.height)
		window.isDirty = true
	})

	window.glw.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		go pointerCoordEvent(&window, x, y)
	})

	return &window
}

func (window *Window) processFrame() {
	if window.isDirty {
		window.isDirty = false
		window.glw.MakeContextCurrent()

		window.surface.SetFillStyle("#FFF")
		window.surface.FillRect(0, 0, float64(window.width), float64(window.height))

		drawRootFrame(window)
		window.glw.SwapBuffers()
	}

	glfw.WaitEvents()
}

//SetRootFrame - Sets the window root frame
func (window *Window) SetRootFrame(frame *Frame) {
	window.rootFrame = frame
}

//Show - Show the window
func (window *Window) Show() {
	window.isDirty = true
	window.visible = true
	window.glw.Show()
}
