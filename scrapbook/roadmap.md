# Roadmap and ideas

- Installing
- Hello World
- Modules
- Packages
- Dependency management
- Testing
- Veting?
- Formatting
- Target platforms, listing and compiling
- Using a Makefile
- Features
  - Functions Functions (are first class citizens)
  - Go Threads
  - Error handling
  - Select?
  - Collections +/-?
  - Interfaces / Empty interface / Monkey typing
  - Generics?
  - Public private
  - Channels
  - Structs
  - Receiver methods
  - Pass by reference / value
- Try something fun
  - Try to generate keys using Elliptic Curve Diffie-Hellman over NIST
  - Play music
  - Slideshow ANSI
  - Better yet, animation on ANSI
  - Karaoke style lyrics text scroll
 
```Text
Go Modules
Go Main
Go Packages
Go Gependency
Go Multithreading
	- Print a banner
	- Play Music
	- Print a timer
Go Testing
Go Compiling
	- Native
	- Linux
	- Windows
	- Mac
	- More: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
Listing Platforms
	- go tool dist list
	- go tool dist list | grep -Po "(?<=/).*$" | sort | uniq
	- go tool dist list | grep -Po "^[^/]*" | sort | uniq
Go Multiple artifacts project
Go Sleep main thread
	- https://stackoverflow.com/questions/36419054/go-projects-main-goroutine-sleep-forever
```

