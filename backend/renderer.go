package backend

import "image"

var platformRenderer Renderer

type Renderer interface {
	Render(*image.RGBA)
}

func init() {
	platformRenderer = createPlatformRenderer()
}

func Render(buffer *image.RGBA) {
	platformRenderer.Render(buffer)
}
