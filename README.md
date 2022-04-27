# Discord Applied Commands
[![Go Report Card](https://goreportcard.com/badge/github.com/vlaetansky/discordac)](https://goreportcard.com/report/github.com/vlaetansky/discordac)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/vlaetansky/discordac?status.svg)](https://godoc.org/github.com/vlaetansky/discordac) 

**Discord Applied Commands** is a Go package that provides an easy way to manage Discord commands.
It provides a high level API to make and register commands for your Discord Bot.

**The package is in an early state of development, any kind of contribution is welcomed!**

## Getting Started
### Installing
`go get github.com/vlaetansky/discordac`

### Usage
Import the package

`import "github.com/vlaetansky/discordac"`

Firstly you need to create a DiscordGo session, it will be referred to as _**discordGoSession**_ further.
Please refer to a DiscordGo documentation to learn on how to create a DiscordGo Session.

Create a new DiscordAC manager which provides a high level API for managing commands

`DAC = discordac.New(discordGoSession)`

Next, the manager must be initialized 

`DAC.Init()`

In order to register application commands you must first open a websocket connection

`discordGoSession.Open()`

You can now create your own commands ([example commands](https://github.com/vlaetansky/discordac/tree/master/examples/dslash/internals/commands)) and register them with

`DAC.RegisterCommands(...SlashedCommands)`

You can register a particular command separately, this can be useful to register a command in a specific guild:

`DAC.RegisterCommand(...SlashedCommands)`

Note: generally you should register your commands using RegisterCommands(..SlashedCommands) method to avoid sending too many command creation requests to Discord API

(Please refer to the [examples](https://github.com/vlaetansky/discordac/tree/master/examples) folder to learn more)
