package pingpong

import (
	"discordslash"
	"github.com/bwmarrin/discordgo"
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
