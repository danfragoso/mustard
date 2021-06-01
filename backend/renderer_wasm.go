package backend

import (
	"fmt"
	"syscall/js"
)

func Render(msg string) {
	fmt.Println("Hi from js")
	renderToCanvas()
	fmt.Println(msg)
}

func renderToCanvas() {
	width := 320
	height := 320

	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	context := canvas.Call("getContext", "2d")

	canvas.Set("height", height)
	canvas.Set("width", width)

	gradient := context.Call("createLinearGradient", 0, 0, 200, 0)
	gradient.Call("addColorStop", 0, "red")
	gradient.Call("addColorStop", 1, "blue")

	context.Set("fillStyle", gradient)
	context.Call("fillRect", 0, 0, 400, 400)
}
