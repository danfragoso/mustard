package mustard

import (
	backend "github.com/danfragoso/mustard/backend"
)

func Render() bool {
	backend.Render("memes")

	return true
}
