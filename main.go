package main

import (
	"github.com/getlantern/systray"
	"strconv"
	"path/filepath"
	"strings"
	popup "github.com/gen2brain/beeep"
	"os"
	"time"
)

func reminderToStayHydrated(interval int, reminderMessage string) {
	timer := time.NewTicker(time.Duration(interval) * time.Minute)
	quit := make(chan struct{})
	go func() {
    	for {
       		select {
        	case <- timer.C:
            	popup.Notify("StayHydrated", reminderMessage, "misc/logo.png")
        	case <- quit:
            	timer.Stop()
            	return
        	}
    	}
 	}()
}

func main() {
	configPath, err := filepath.Abs("StayHydrated.ini")
	if err != nil {
		os.Exit(1)
	}

	if !ConfigExists(configPath) {
		CreateBlankConfig(configPath)
		popup.Notify("StayHydrated", "Config file created at " + configPath + "\nYou can modify it later!", "misc/logo.png")
	} else {
		popup.Notify("StayHydrated", "Succesfully read config!", "misc/logo.png")
		arguments := ReadConfig(configPath)
		interval, _ := strconv.Atoi(strings.Split(arguments[0], "min")[0])
		customMessage := arguments[1]

		if arguments[2] == "true" {
			EnableAutostart()
		}

		if arguments[2] == "false" {
			DisableAutostart()
		}
		
		reminderToStayHydrated(interval, customMessage)
	}
	systray.Run(RunUI, QuitUI)
}