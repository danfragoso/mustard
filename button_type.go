package mustard

//ButtonClickCallback - The click event callback
type ButtonClickCallback func(*ButtonWidget)

//ButtonWidget - Button Widget
type ButtonWidget struct {
	widget

	content string

	onClick ButtonClickCallback
}
