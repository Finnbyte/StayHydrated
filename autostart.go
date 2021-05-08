package main

import (
	"github.com/emersion/go-autostart"
	"path/filepath"
	"os"
)

func EnableAutostart() {
	fpath, err := filepath.Abs("./StayHydrated.exe")
	if err != nil {
		
	}
	app := &autostart.App{
		Name: "StayHydrated",
		DisplayName: "Water Drinking Reminder app.",
		Exec: []string{"start", fpath},
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
	fpath, err := filepath.Abs("./StayHydrated.exe")
	if err != nil {
		
	}
	app := &autostart.App{
		Name: "StayHydrated",
		DisplayName: "Water Drinking Reminder app.",
		Exec: []string{"start", fpath},
	}

	if app.IsEnabled() {
		if err := app.Disable(); err != nil {
			os.Exit(1)
		}
	} else {
		return
	}
}