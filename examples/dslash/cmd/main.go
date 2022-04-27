package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/vlaetansky/discordslash"
	"github.com/vlaetansky/discordslash/examples/dslash/internals/commands/pingpong"
	"github.com/vlaetansky/discordslash/examples/dslash/internals/commands/wikipedia"
	"log"
	"os"
	"os/signal"
)

const (
	TokenEnv string = "DGO_TOKEN"
	GIdEnv   string = "GUILD_ID"
)

var (
	token   string
	guildId string
	DAC     *discordac.DiscordAC
)

func init() {
	token = os.Getenv(TokenEnv)

	if token == "" {
		panic("DGO_TOKEN env. variable is not specified")
	}

	guildId = os.Getenv(GIdEnv)

	if guildId == "" {
		panic("GUILD_ID env. variable is not specified")
	}
}

func main() {
	session, err := discordgo.New(fmt.Sprintf("Bot %v", token))
	if err != nil {
		return
	}

	DAC = discordac.New(session)
	DAC.Init()

	err = session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	logrus.Info("Registering commands...")
	err = DAC.RegisterCommands(guildId,
		pingpong.Command,
		wikipedia.Command,
	)

	if err != nil {
		logrus.WithError(err).Error("Couldn't register commands")
	}

	defer func(session *discordgo.Session) {
		err := session.Close()
		if err != nil {
			logrus.WithError(err).Info("Couldn't properly close websocket connection to Discord")
		}
		logrus.Info("Bye!")
	}(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logrus.Info("Press Ctrl+C to exit")
	<-stop
}
