package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	app := app.New()
	window := app.NewWindow("StayHydrated")

	hello := widget.NewLabel("Hello! Just keep this program on, and it will remind you to drink water! :)")
	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Epic.", func() {
			return
		}),
	))

	window.ShowAndRun()
}