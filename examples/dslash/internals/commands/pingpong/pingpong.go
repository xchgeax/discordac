package pingpong

import (
	"github.com/bwmarrin/discordgo"
	"github.com/vlaetansky/discordslash"
)

var Command = discordslash.SlashedCommand{
	Specification: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping Pong Command",
	},
	Handler: func(cc *discordslash.CommandContext) {
		cc.Respond("Pong!")
	},
}
