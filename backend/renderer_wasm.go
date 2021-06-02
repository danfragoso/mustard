package backend

import (
	"fmt"
	"image"
	"syscall/js"

	"github.com/danfragoso/mustard/device"
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
	w, h := device.GetViewportDimensions()

	canvas.Set("width", int(w))
	canvas.Set("height", int(h))

	context := canvas.Call("getContext", "2d")
	return &wasmRenderer{
		msg: "Hi from the wasm renderer",

		canvas:  canvas,
		context: context,
	}
}

func renderToCanvas(buffer *image.RGBA, canvas, context js.Value) {
	js.CopyBytesToJS(js.Global().Get("pixData"), buffer.Pix)
	js.Global().Call("renderNative", buffer.Stride)
}
