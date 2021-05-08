package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	app := app.New()
	window := app.NewWindow("StayHydrated")

	widget.NewLabel("Hello! Just keep this program on, and every x minutes it will remind you to drink water! :)")
	window.ShowAndRun()
}