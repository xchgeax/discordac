package discordslash

import "github.com/bwmarrin/discordgo"

// SlashedCommand is a command representation inside DiscordSlash
// Wraps command specification described by discordgo.ApplicationCommand
// and holds a corresponding command handler
type SlashedCommand struct {
	Specification *discordgo.ApplicationCommand
	// TODO: find a better approach with guild id?
	GuildId string
	// Command implementation
	Handler func(cc *CommandContext)
}

// Global determines either SlashCommand is defined to run globally or in SlashedCommand.GuildId
func (sc SlashedCommand) Global() bool {
	return sc.GuildId == ""
}

// Invoke calls SlashedCommand.Handler
// cc	: CommandContext around which the command will run
func (sc SlashedCommand) Invoke(cc *CommandContext) {
	cc.parseOptions()
	sc.Handler(cc)
}

// Name returns command name
func (sc SlashedCommand) Name() string {
	return sc.Specification.Name
}
