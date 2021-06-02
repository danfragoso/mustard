package mustard

import (
	"image"
	"image/color"

	backend "github.com/danfragoso/mustard/backend"
	device "github.com/danfragoso/mustard/device"
)

func Render() bool {
	w, h := device.GetViewportDimensions()
	backend.Render(imgToRender(w, h))

	return true
}

func imgToRender(w, h float64) *image.RGBA {
	width := int(w)
	height := int(h)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	green := color.RGBA{10, 200, 100, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, green)
			default:
				img.Set(x, y, color.Black)
			}
		}
	}

	return img
}
