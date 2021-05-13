package main

import (
	"os"
	"github.com/emersion/go-autostart"
)

func EnableAutostart() {
	app := &autostart.App{
		Name: "StayHydrated",
		DisplayName: "Water Drinking Reminder app.",
	}

	if app.IsEnabled() {
		return
	} else {
		if err := app.Enable(); err != nil {
			os.Exit(1)
		}
	}
}

func DisableAutostart() {
	app := &autostart.App{
		Name: "StayHydrated",
		DisplayName: "Water Drinking Reminder app.",
	}

	if app.IsEnabled() {
		if err := app.Disable(); err != nil {
			os.Exit(1)
		}
	} else {
		return
	}
}