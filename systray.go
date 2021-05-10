package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func RunUI() {
	systray.SetIcon(icon.Data)
	systray.SetTooltip("StayHydrated")
	quitButton := systray.AddMenuItem("Quit", "Quit StayHydrated")

	go func() {
		<-quitButton.ClickedCh
		systray.Quit()
	}()

}

func QuitUI() {

}