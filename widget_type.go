package mustard

type WidgetOrientation int

const (
	RowWidgetOrientation WidgetOrientation = iota
	ColumnWidgetOrientation
)

type Widget struct {
	top    float64
	left   float64
	width  float64
	height float64

	dirty bool

	widgets []interface{}

	backgroundColor string
	orientation     WidgetOrientation
}
