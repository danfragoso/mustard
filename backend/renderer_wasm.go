package backend

import (
	"fmt"
	"image"
	"syscall/js"
)

type wasmRenderer struct {
	msg string

	canvas  js.Value
	context js.Value
}

func (r *wasmRenderer) Render(buffer *image.RGBA) {
	fmt.Println(r.msg)
	renderToCanvas(buffer, r.canvas, r.context)
}

func createPlatformRenderer() Renderer {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	canvas.Set("height", 400)
	canvas.Set("width", 400)
	context := canvas.Call("getContext", "2d")
	return &wasmRenderer{
		msg: "Hi from the wasm renderer",

		canvas:  canvas,
		context: context,
	}
}

func renderToCanvas(buffer *image.RGBA, canvas, context js.Value) {
	width := 400
	height := 400

	js.CopyBytesToJS(js.Global().Get("pixData"), buffer.Pix)
	js.Global().Call("renderNative", buffer.Stride)
	js.Global().Get("console").Call("time", "wasm")

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := buffer.RGBAAt(x, y)
			js.Global().Call("putPixelData", x, y, c.R, c.G, c.B, c.A)
		}
	}

	js.Global().Get("console").Call("timeEnd", "wasm")
}
