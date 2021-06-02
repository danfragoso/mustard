package backend

import (
	"fmt"
	"image"
)

type androidRenderer struct {
	msg string
}

func (r *androidRenderer) Render(buffer *image.RGBA) {
	fmt.Println(r.msg)
}

func createPlatformRenderer() Renderer {
	return &androidRenderer{
		msg: "Hi from the android renderer",
	}
}
