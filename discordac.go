package discordac

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// DiscordAC represents a high level API to command management
type DiscordAC struct {
	session *discordgo.Session
}

// New creates a new DiscordAC manager
// s	: bot session
func New(s *discordgo.Session) (dislash *DiscordAC) {
	dislash = &DiscordAC{session: s}

	return
}

// Init initializes command dispatcher by adding a new handler which dispatches and invokes commands
func (ds *DiscordAC) Init() {
	ds.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if command, ok := dispatchCommand(i.ApplicationCommandData().ID); ok {
			go command.Invoke(&CommandContext{Session: s, Interaction: i.Interaction})
		} else {
			logrus.Errorf("Coudn't dispatch command %v: no handler registered", i.ApplicationCommandData().Name)
		}
	})
}

// RegisterCommands registers a list of commands
// guildId  : guild to register the commands in, leave blank to register it globally
// commands	: commands to register
func (ds *DiscordAC) RegisterCommands(guildId string, commands ...*AppliedCommand) error {
	var specifications []*discordgo.ApplicationCommand

	for _, command := range commands {
		if command.Specification.Type == 0 {
			panic(fmt.Sprintf("Command type must be specified for a command %v", command.Name()))
		}
		specifications = append(specifications, command.Specification)
	}

	createdCommandsSpecs, err := ds.session.ApplicationCommandBulkOverwrite(ds.session.State.User.ID, guildId, specifications)
	if err != nil {
		return err
	}

	// Update commands specifications
	// Discord returns a modified specification with an assigned command ID
	for _, newSpecification := range createdCommandsSpecs {
		for _, command := range commands {
			if newSpecification.Type == command.Specification.Type && newSpecification.Name == command.Name() {
				command.Specification = newSpecification
				command.guildId = guildId
			}
		}
	}

	addCommands(commands...)

	return nil
}

// RegisterCommand registers a single command
// guildId  : guild to register the commands in, leave blank to register it globally
// command	: command to register
func (ds *DiscordAC) RegisterCommand(guildId string, command *AppliedCommand) error {
	if command.Specification.Type == 0 {
		panic(fmt.Sprintf("Command type must be specified for a command %v", command.Name()))
	}

	createdCommandSpec, err := ds.session.ApplicationCommandCreate(ds.session.State.User.ID, guildId, command.Specification)
	if err != nil {
		return err
	}

	command.Specification = createdCommandSpec
	command.guildId = guildId

	addCommands(command)

	return nil
}

// UnregisterCommands unregisters all commands both globally and per guild
// This should be called during bot shutdown
func (ds *DiscordAC) UnregisterCommands() {
	for _, command := range registeredCommands {
		err := ds.session.ApplicationCommandDelete(ds.session.State.User.ID, command.guildId, command.Specification.ID)
		if err != nil {
			logrus.WithError(err).Panicf("Couldn't unregister '%v' command", command.Specification.Name)
		}
	}
}
