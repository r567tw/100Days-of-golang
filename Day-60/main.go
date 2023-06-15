package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("First App")

	hello := widget.NewLabel("Hello World")
	input := widget.NewEntry()

	w.SetContent(container.NewVBox(
		hello,
		input,
		widget.NewButton("Hi!", func() {
			hello.SetText(fmt.Sprintf("Welcome %s :)", input.Text))
		}),
	))

	w.ShowAndRun()
}
