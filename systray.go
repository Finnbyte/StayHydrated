package main

import (
	"os/exec"
	"github.com/getlantern/systray"
	"./utils"
)

func RunUI() {
	systray.SetIcon(utils.LogoData)
	systray.SetTooltip("StayHydrated")
	infoButton := systray.AddMenuItem("About", "Open StayHydrated's Github repository.")
	quitButton := systray.AddMenuItem("Quit", "Quit StayHydrated.")

	go func() {
		<-quitButton.ClickedCh
		systray.Quit()
		
	}()

	go func() {
		<-infoButton.ClickedCh
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/Finnbyte/StayHydrated").Start()
	}()

}

func QuitUI() {

}