package main

import (
	"strconv"
	"path/filepath"
	"strings"
	popup "github.com/gen2brain/beeep"
	"os"
	"time"
)

func reminderToStayHydrated(interval int) {
	timer := time.NewTicker(time.Duration(interval) * time.Minute)
	quit := make(chan struct{})
	go func() {
    	for {
       		select {
        	case <- timer.C:
            	popup.Notify("StayHydrated", "It's time to get sum hydration!!!", "")
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
		popup.Notify("StayHydrated", "Config file created at " + configPath, "")
		os.Exit(1)
	} else {
		popup.Notify("StayHydrated", "Succesfully read config!", "")
		arguments := ReadConfig(configPath)

		autoStart := arguments[1]

		if autoStart {
			EnableAutostart()
		} else {
			DisableAutostart()
		}

		// TIME IN MINUTES
		// Remove "min" because it's overcomplicating things and I don't see a point
		// TODO: ADD OPTION IN UI TO MODIFY INTERVAL

		interval := strconv.Atoi(arguments[0]) 

		reminderToStayHydrated(interval)
	}

	RunUI()
}