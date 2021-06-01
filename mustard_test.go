package mustard

import (
	"fmt"
	"testing"
)

func TestRender(t *testing.T) {
	out := Render()
	fmt.Println(out)
}
