package main

import (
	"fmt"

	"github.com/danfragoso/mustard"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/touch"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case touch.Event:
				fmt.Println("MUSTARD EV:", e)
			case lifecycle.Event:
				fmt.Println("MUSTARD EV:", e)
			case paint.Event:
				fmt.Println("MUSTARD EV: Render???")
				mustard.Render()
				a.Publish()
			}
		}
	})
}
