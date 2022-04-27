package discordslash

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// DiscordSlash represents a high level API to command management
type DiscordSlash struct {
	session *discordgo.Session
}

// New creates a new DiscordSlash manager
// s	: bot session
func New(s *discordgo.Session) (dislash *DiscordSlash) {
	dislash = &DiscordSlash{session: s}

	return
}

// Init initializes command dispatcher by adding a new handler which dispatches and invokes commands
func (ds *DiscordSlash) Init() {
	ds.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		requestedInteraction := i.ApplicationCommandData().Name
		if command, ok := dispatchCommand(requestedInteraction); ok {
			go command.Invoke(&CommandContext{Session: s, Interaction: i.Interaction})
		} else {
			logrus.Errorf("Coudn't dispatch command %v: no handler registered", requestedInteraction)
		}
	})
}

// RegisterCommands registers a list of commands
// guildId  : guild to register the commands in, leave blank to register it globally
// commands	: commands to register
func (ds *DiscordSlash) RegisterCommands(guildId string, commands ...*SlashedCommand) error {
	var specifications []*discordgo.ApplicationCommand

	// Save specifications of provided commands and check if they are unique
	for _, command := range commands {
		if commandExist(command.Name()) {
			panic(fmt.Sprintf("Command already registered: %v", command.Name()))
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
			}
		}
	}

	addCommands(commands...)

	return nil
}

// RegisterCommand registers a single command
// guildId  : guild to register the commands in, leave blank to register it globally
// command	: command to register
func (ds *DiscordSlash) RegisterCommand(guildId string, command *SlashedCommand) error {
	if commandExist(command.Name()) {
		panic(fmt.Sprintf("Command already registered: %v", command.Name()))
	}

	createdCommandSpec, err := ds.session.ApplicationCommandCreate(ds.session.State.User.ID, guildId, command.Specification)
	if err != nil {
		return err
	}

	command.Specification = createdCommandSpec

	addCommands(command)

	return nil
}

// UnregisterCommands unregisters all commands both globally and per guild
// This should be called during bot shutdown
func (ds *DiscordSlash) UnregisterCommands() {
	for _, command := range registeredCommands {
		err := ds.session.ApplicationCommandDelete(ds.session.State.User.ID, command.GuildId, command.Specification.ID)
		if err != nil {
			logrus.WithError(err).Panicf("Couldn't unregister '%v' command", command.Specification.Name)
		}
	}
}
