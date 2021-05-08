package main

import (
	ini "gopkg.in/ini.v1"
	"os"
	popup "github.com/gen2brain/beeep"
)

func ConfigExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil { //if config does not exist, return false
		return false
	}
	return true
}

func CreateBlankConfig(path string) {
	_, err := os.Create(path)
	if err != nil {
		os.Exit(1)
	}

	cfg, errr := ini.Load(path)
	if errr != nil {
		popup.Notify("StayHydrated", "Error occured!", "")
		os.Exit(1)
	}

	cfg.NewSection("Settings")
	cfg.Section("Settings").NewKey("Reminder Interval", "x")
	cfg.Section("Settings").NewKey("Autostart With Windows", "x")

	cfg.SaveTo(path)
}

func ReadConfig(path string) []string {
	var config []string
	cfg, err := ini.Load(path)
	if err != nil {
		popup.Notify("StayHydrated", "Error occured!", "")
	}

	interval, _ := cfg.Section("Settings").GetKey("Reminder Interval")
	autostarting, _ := cfg.Section("Settings").GetKey("Autostart With Windows")

	config = append(config, interval.Value())
	config = append(config, string(autostarting.Value()))

	return config
} 