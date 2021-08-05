package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func NewContentRect(size uint) *canvas.Rectangle {
	rect := canvas.NewRectangle(theme.BackgroundColor())
	rect.StrokeColor = theme.ShadowColor()
	rect.StrokeWidth = theme.SeparatorThicknessSize()
	rect.SetMinSize(fyne.NewSize(float32(size), 1))
	return rect
}
