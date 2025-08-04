package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type emptyTheme struct {
}

func (e *emptyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return nil
}

func (e *emptyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return nil
}

func (e *emptyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return nil
}

func (e *emptyTheme) Size(n fyne.ThemeSizeName) float32 {
	return 0
}
func main() {
	a := app.New()
	a.Settings().SetTheme(&emptyTheme{})
	w := a.NewWindow("Hello")
	var count int
	hello := widget.NewLabel(fmt.Sprintf("current count:\t%d", count))
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Increment Count", func() {
			count++
			hello.SetText(fmt.Sprintf("current count:\t%d", count))
		}),
	))

	w.ShowAndRun()
}
