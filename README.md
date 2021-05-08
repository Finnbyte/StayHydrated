# StayHydrated
Incredibly simple yet lifesaving app with a simple goal: GET YOU TO DRINK THE MOST VALUABLE LIQUID IN THE UNIVERSE.

# THE GUI IS VERY SLOW RN, WILL FIX SOON. JUST WAIT A BIT FOR IT TO POP OPEN!

## Requirements
None if using binary version, but if you want to compile from src, you will need Golang installed (v.1.15) and you also need to download the dependencies. I assume you know what you are doing.

## How to use
1. First, compile using
`go build -ldflags -H=windowsgui -o StayHydrated.exe main.go autostart.go config.go ui.go` or just download from releases.
2. Run StayHydrated.exe
3. It will pop up a notification that config file has been created.
4. Open config file with a text editor.
5. Refer to Config File Guide for usage
6. Launch the app again and now it will prompt you to drink water every specified amount of minutes!

## Config File Guide
It's very simple to put correct values to the config file. 
Reminder Interval needs to be a number and needs to end with min (eg. 40min)
Autostarting With Windows can only be **true** or **false**
