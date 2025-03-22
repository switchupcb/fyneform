package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/fyneform/example/app/fyneform"
)

func main() {
	app := app.New()

	window := app.NewWindow("fyneform")

	form := fyneform.AccountForm()

	window.SetContent(form)
	window.Resize(fyne.NewSize(form.Size().Width+64, form.Size().Height))
	window.ShowAndRun()
}
