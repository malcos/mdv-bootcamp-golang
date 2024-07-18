# Hey ho! Let's Go!

Go is a statically typed, compiled high-level programming language.

It is syntactically similar to C, but also provides memory safety, garbage collection, structural typing, lightweight green threads and CSP-style concurrency.

It is able to compile native executables for different operating systems and hardware platforms.

## Installing

You can either [download and install](https://go.dev/doc/install) Go manually, or use the package manager of your OS.

Windows:

```shell
choco install golang
```

Darwin:

```shell
brew install go
```

Linux:

```shell
# C'mon. Use you package manager, as usual.
```

Verify your installed version:

```shell
$ go version
go version go1.22.5 linux/amd64
```



## Useful resources and links

- Learning go: https://go.dev/learn/
  - Interactive tour of Go: https://go.dev/tour
  - Go by Example: https://gobyexample.com/
- Documentation: https://go.dev/doc/
  - Installing: https://go.dev/doc/install
  - Effective Go: https://go.dev/doc/effective_go
  - The Standard library: https://pkg.go.dev/std
- Social contacts and Help: https://go.dev/help
  - The Go Blog: https://go.dev/blog/
  - Conferences and Events: https://go.dev/wiki/Conferences
  - Meetup: https://www.meetup.com/pro/go/
- Release notes: https://go.dev/doc/devel/release
- ~~Maven Central~~ Go Packages: https://pkg.go.dev/

## TOPICS / TODO

The list is incomplete and requires reordering

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
 
```Tex
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

