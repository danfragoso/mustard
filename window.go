package mustard

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
	"github.com/veandco/go-sdl2/sdl"
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

	sdlw, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(width), int32(height), sdl.WINDOW_RESIZABLE|sdl.WINDOW_OPENGL|sdl.WINDOW_HIDDEN)
	if err != nil {
		fmt.Printf("Error creating window: %v", err)
	}

	sdlwID, err := sdlw.GetID()
	if err != nil {
		fmt.Printf("Error getting window ID: %v", err)
	}

	glContext, err := sdlw.GLCreateContext()
	if err != nil {
		fmt.Printf("Error creating GL context: %v", err)
	}

	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}

	sdl.GLSetSwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	window.sdlw = sdlw
	window.sdlwID = sdlwID
	window.backend = backend
	window.glContext = glContext
	window.isDirty = true
	window.surface = canvas.New(backend)

	return &window
}

func (window *Window) processFrame() {
	window.fireEvents()
	window.updateSize()

	if window.isDirty {
		window.sdlw.GLMakeCurrent(window.glContext)

		window.backend.SetBounds(0, 0, window.width, window.height)

		window.surface.SetFillStyle("#FFF")
		window.surface.FillRect(0, 0, float64(window.width), float64(window.height))

		drawRootFrame(window)

		if window.frames > 0 {
			window.isDirty = false
		}

		window.sdlw.GLSwap()
		window.frames = window.frames + 1
	}
}

//SetRootFrame - Sets the window root frame
func (window *Window) SetRootFrame(frame *Frame) {
	window.rootFrame = frame
}

//Show - Show the window
func (window *Window) Show() {
	window.sdlw.Show()
	window.visible = true
}

//Hide - Hide the window
func (window *Window) Hide() {
	window.sdlw.Hide()
	window.visible = false
}

//updateSize - Update the window size
func (window *Window) updateSize() {
	w, h := window.sdlw.GetSize()
	width, height := int(w), int(h)

	if window.width != width || window.height != height {
		window.width, window.height = width, height
		window.isDirty = true
	}
}

//fireEvents - Poll sdl events and fire them
func (window *Window) fireEvents() {
	for {
		event := sdl.PollEvent()

		if event == nil {
			break
		}
	}
}
