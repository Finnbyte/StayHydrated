package main

import (
	ini "gopkg.in/ini.v1"
	"os"
	popup "github.com/gen2brain/beeep"
)

func ConfigExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func CreateBlankConfig(path string) {
	_, err := os.Create(path)
	if err != nil {
		os.Exit(1)
	}

	cfg, err := ini.Load(path) // Reuse err variable as it doesnt seem to be a const variable?
	if err != nil {
		popup.Notify("StayHydrated", "Error occured!", "")
		os.Exit(1)
	}

	// DEFAULT SETTINGS
	cfg.NewSection("Settings")
	cfg.Section("Settings").NewKey("Reminder Interval", "40")
	cfg.Section("Settings").NewKey("Autostart With Windows", "false")

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
	config = append(config, string(autostarting.Value())) // why are you turning a string into a string?

	return config
} 