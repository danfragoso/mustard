package mustard

type widget struct {
	top    int
	left   int
	width  int
	height int

	dirty       bool
	fixedWidth  bool
	fixedHeight bool

	widgets []interface{}

	backgroundColor string
}
