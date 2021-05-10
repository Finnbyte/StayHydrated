package main

import (
	"strings"
	"strconv"
	ini "gopkg.in/ini.v1"
	"os"
	popup "github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
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
		popup.Notify("StayHydrated", "Error occured!", "misc/logo.png")
		os.Exit(1)
	}

	// ask user about options

	dlgs.Info("StayHydrated", "StayHydrated will ask you some questions now. Answer based on your preferences!")
	intervalString, _, _ := dlgs.Entry("StayHydrated", "How often reminded?", "")
	autostartBool, _ := dlgs.Question("StayHydrated", "Do you want StayHydrated to launch automatically when turning on your computer?", true)

	cfg.NewSection("Settings")
	cfg.Section("Settings").NewKey("ReminderInterval", intervalString)
	cfg.Section("Settings").NewKey("AutostartWithWindows", strconv.FormatBool(autostartBool))
	cfg.SaveTo(path)

	interval, _ := strconv.Atoi(strings.Split(intervalString, "min")[0])
	if err != nil {
		popup.Notify("StayHydrated", "Bad interval value!", "misc/logo.png")
	}
	reminderToStayHydrated(interval)
}

func ReadConfig(path string) []string {
	var config []string
	cfg, err := ini.Load(path)
	if err != nil {
		popup.Notify("StayHydrated", "Error occured!", "")
	}

	interval, _ := cfg.Section("Settings").GetKey("ReminderInterval")
	autostarting, _ := cfg.Section("Settings").GetKey("AutostartWithWindows")

	config = append(config, interval.Value())
	config = append(config, string(autostarting.Value()))

	return config
} 