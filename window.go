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

	window.Visible = false
	window.Initialized = false

	window.Title = windowTitle

	window.Width = width
	window.Height = height

	window.Top = top
	window.Left = left

	glw, err := glfw.CreateWindow(window.Width, window.Height, window.Title, nil, nil)
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

	return &window
}

func (window *Window) loop() {
	for !window.glw.ShouldClose() {
		window.Width, window.Height = window.glw.GetSize()

		if window.isDirty {
			window.glw.MakeContextCurrent()
			window.backend.SetBounds(0, 0, window.Width, window.Height)
			window.surface.SetFillStyle("#FFF")
			window.surface.FillRect(0, 0, float64(window.Width), float64(window.Height))

			drawWidgets(window.Widgets, window.surface, 0, 0)
			window.glw.SwapBuffers()
		}

		glfw.WaitEvents()
	}
}

//AttachWidget - Attaches a new widget to the window
func (window *Window) AttachWidget(widget interface{}) {
	window.Widgets = append(window.Widgets, widget)
}

//Show - Show the window
func (window *Window) Show() {
	window.isDirty = true
	window.glw.Show()
	go window.loop()
}
