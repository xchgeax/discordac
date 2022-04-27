package discordac

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// AppliedCommand is a command representation inside DiscordAC
// Wraps command specification described by discordgo.ApplicationCommand
// and holds a corresponding command handler
type AppliedCommand struct {
	Specification *discordgo.ApplicationCommand
	// TODO: find a better approach with guild id?
	GuildId string
	// Command implementation
	Handler func(cc *CommandContext)
}

// Global determines either SlashCommand is defined to run globally or in AppliedCommand.GuildId
func (sc AppliedCommand) Global() bool {
	return sc.GuildId == ""
}

func (sc AppliedCommand) InternalName() string {
	return fmt.Sprintf("type%v_%v", sc.Specification.Type, sc.Name())
}

// Invoke calls AppliedCommand.Handler
// cc	: CommandContext around which the command will run
func (sc AppliedCommand) Invoke(cc *CommandContext) {
	cc.parseOptions()
	sc.Handler(cc)
}

// Name returns command name
func (sc AppliedCommand) Name() string {
	return sc.Specification.Name
}
