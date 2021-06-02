package device

import "syscall/js"

func getViewportDimensions() (float64, float64) {
	width := js.Global().Get("innerWidth").Float()
	height := js.Global().Get("innerHeight").Float()

	return width, height
}
