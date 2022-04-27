package pingpong

import (
	"github.com/bwmarrin/discordgo"
	"github.com/vlaetansky/discordslash"
)

var Command = &discordac.AppliedCommand{
	Specification: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping Pong Command",
		Type:        discordgo.ChatApplicationCommand,
	},
	Handler: func(cc *discordac.CommandContext) {
		cc.Respond("Pong!")
	},
}
