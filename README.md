# Discord Slash

**Discord Slash** is a Go package that provides an easy way to manage Discord's new slash commands.
It provides a high level API to make and register slash commands for your Discord Bot.

**The package is in an early state of development, any kind of contribution is welcomed!**

## Getting Started
### Installing
`go get github.com/vlaetansky/discordslash`

### Usage
Import the package

`import "github.com/vlaetansky/discordslash"`

Please refer to a DiscordGo documentation to read on how to create a DiscordGo Session.

Create a new DiscordSlash manager which provides a high level API for managing commands

`DiSlash = discordslash.New(discordGoSession)`

Next, the manager must be initialized 

`DiSlash.Init()`

You can now create your own commands and register them with one of the options

`DiSlash.RegisterCommand(SlashedCommand)`

`DiSlash.RegisterCommands(...SlashedCommands)`

`DiSlash.RegisterCommandWithin(guildId, SlashedCommand)`

`DiSlash.RegisterCommandsWithin(guildId, ...SlashedCommands)`

During bot shutdown, you should also unregister all the commands

`DiSlash.UnregisterCommands()`

(Please refer to the [examples](https://github.com/vlaetansky/discordslash/tree/master/examples) folder to learn how to create commands)
