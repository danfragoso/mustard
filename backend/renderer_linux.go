// +build !android

package backend

import (
	"fmt"
	"image"
)

type linuxRenderer struct {
	msg string
}

func (r *linuxRenderer) Render(buffer *image.RGBA) {
	fmt.Println(r.msg)
}

func createPlatformRenderer() Renderer {
	return &linuxRenderer{
		msg: "Hi from the linux renderer",
	}
}
