package discordslash

import (
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

// RegisterCommandWithin registers a command for a specific guild
// command	: command to register
// guildId	: ID of a guild to register the command for
func (ds *DiscordSlash) RegisterCommandWithin(command SlashedCommand, guildId string) (err error) {
	command.GuildId = guildId
	err = ds.RegisterCommand(command)

	return
}

// RegisterCommandsWithin registers a list of commands for a specific guild
// guildId		: ID of a guild to register commands for
// commands		: commands to register
func (ds *DiscordSlash) RegisterCommandsWithin(guildId string, commands ...SlashedCommand) {
	for _, command := range commands {
		err := ds.RegisterCommandWithin(command, guildId)
		if err != nil {
			logrus.WithError(err).Errorf("Couldn't register command %v within %v", command.Name(), guildId)
		}
	}
}

// RegisterCommands registers a list of commands globally
// commands	: commands to register
func (ds *DiscordSlash) RegisterCommands(commands ...SlashedCommand) {
	for _, command := range commands {
		err := ds.RegisterCommand(command)
		if err != nil {
			logrus.WithError(err).Errorf("Couldn't register command %v", command.Name())
		}
	}
}

// RegisterCommand registers a command globally
// command	: command to register
func (ds *DiscordSlash) RegisterCommand(command SlashedCommand) (err error) {
	if commandExist(command.Name()) {
		logrus.Panicf("Command already registered: %v", command.Name())
	}

	if command.Global() {
		logrus.Infof("Registering new global command %v...", command.Name())
	} else {
		logrus.Infof("Registering new command %v within %v", command.Name(), command.GuildId)
	}

	// Tell discord about our commands specifications
	command.Specification, err = ds.session.ApplicationCommandCreate(ds.session.State.User.ID, command.GuildId, command.Specification)
	if err != nil {
		return err
	}
	// Add a command to the dispatcher list to invoke and unregister it later
	addCommand(command)

	return nil
}

// UnregisterCommands unregisters all commands both globally and per guild
// This should be called during bot shutdown
func (ds *DiscordSlash) UnregisterCommands() {
	logrus.Info("Unregistering commands...")
	for _, command := range commands {
		err := ds.session.ApplicationCommandDelete(ds.session.State.User.ID, command.GuildId, command.Specification.ID)
		if err != nil {
			logrus.WithError(err).Panicf("Couldn't unregister '%v' command", command.Specification.Name)
		}
	}
}
