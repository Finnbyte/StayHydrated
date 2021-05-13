<img src="https://github.com/Finnbyte/StayHydrated/blob/master/misc/logo.png" width="500" height="500">

# StayHydrated
Incredibly simple yet essential app with a simple goal: GET YOU TO DRINK THE MOST VALUABLE LIQUID IN THE UNIVERSE.

## Why I created this
Did you know that 75% of americans are constantly dehydrated? That's scary. Dehydration causes fatigue, dizziness and it lowers your mental and physical performance. That performance is especially important for people like me, who love programming and gaming. It can make you much worse at performing in games and can also make or break a session. **I definitely have chronic dehydration and I want to help people who are cursed by it too!**

## Installation
There are two methods of obtaining this piece of software:
1. You can manually compile. I will be assuming you know what to do and I won't be covering this entire process here, but essentially, first off you git clone this repo and then ***YOU MUST COMPILE*** with `go build -ldflags -H="windowsgui" main.go config.go autostart.go systray.go` after you `go get` all 3rd party dependecies, obviously.
2. The most simple way is to go to https://github.com/Finnbyte/StayHydrated/releases and download StayHydrated.exe from there.

## USAGE
I wanted to make usability as simple as possible. I definitely think this is an important app for anyone, and it was important to me that my friends and family who are pretty dumb in things like this, could just easily go to this site, download this app and be easily able to integrate it into their lives. However, full steps to getting this running are still here just in case! :)
1. You download/compile this to your computer.
2. Run the app.
3. It will now open up a few windows where you need to input some needed things eg. reminder message
4. Once that's done, it will now live in your systray
<img src="https://user-images.githubusercontent.com/58516858/118126471-137f2a00-b401-11eb-9e10-7affbf71f450.png" width="100" height="50">

### IF AFTER ANSWERING TO THE TEXTBOXES THE PROGRAM DOES NOT OPEN, CHECK YOUR CONFIG FILE FOR TYPOS!!! Check CONFIGURING for more information
