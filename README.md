# chatgo

A WebSockets based chat application, written in Go

(If you want to know more about the project specifically, skip to the **Project** section)

## About

After a long time of switching between technologies, I decided that I wanted to start poking around in backend development, because it honestly sounded the most fun to me. After thinking about which language I wanted to start learning backend for, I decided on Go.

Why Go?

From what I've seen, people who use Go, either for backend development or something else, they've only had good things to say about it, so I wanted to play with it too. (I also wanted to try out a functional language, after using only OOP languages before this, and Go felt like a good middle ground)

After fumbling around with how Go works, and figuring out some of the things that make it different, by making little mini-projects here and there, I finally decided to do what Prime said-build a WebSockets chat app.

So here we are, with a (not very featured, I'll admit) chat app, written in Go with gorilla/websocket.

## Project

**PLEASE NOTE: FUNCTIONALITY OF THIS APPLCATION IS VERY LIMITED (2025-01-08)**

`chatgo` is a CLI-Based, WebSocket chat app written in go.

As of now, this is VERY, VERY WIP. However, you can:
* Start a server on localhost: `chatgo server`
* Connect clients to said server: `chatgo client <username>`

The server will log all messages sent in `$APPLICATIONPATH/data/timestamp.txt`

(Note that if you change the URL that the server is serving to, you MUST change it for the client-side in the source code)

I plan to continue working on this project, as it has actually been very, very fun for me!

Some of the features that I plan on adding are:
* A [bubbletea](https://github.com/charmbracelet/bubbletea) based frontend client (the bones of this are starting in [src/model](https://github.com/Gammer0909/chatgo))
* A server customizing system, probably through yaml
* A more intuitive and probably better parsed CLI system
* More basic chat application functionality (eg. whispering, sign-ins, etc)

If you have any features you'd like to see, make an issue and we'll talk about it!

One of my main ideas for this project is that due to it being open-sourced, and you setup the server yourself, I want this application to be mega-customizable, so ideally I'd like:
* Custom Clients (through whatever language or library that supports websockets)
* Custom Servers (not just customization through yaml files, purely custom server handling)

If there's anything else you'd like to see customizeable, again, make an issue and we'll talk about that!

## Building

If you don't feel like building this application, please grab the newest executable from the [releases page](https://github.com/Gammer0909/chatgo/releases)!

If you *do* feel like building the application, first make sure you have go 1.23.4 installed.

Then, grab the source code:
```
git clone https://github.com/Gammer0909/chatgo.git
# Or, using SSH:
git clone git@github.com:Gammer0909/chatgo.git
```

Make sure you download the dependencies:
```
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/gorilla/websocket
go get github.com/charmbracelet/bubbles
```
Then, everything should be ready to build!
`go build -o chatgo`

## Contributing

Please don't
im not ready yet

# License

This project is under the MIT license.
