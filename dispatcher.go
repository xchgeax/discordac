package discordac

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// commands is a list of registered AppliedCommand
	registeredCommands = make(map[string]*AppliedCommand)
)

type DispatchRequest struct {
	Name string
	Type discordgo.InteractionType
}

// addCommands adds commands to the dispatcher list
// to dispatch and unregister it later
// commands: command list to register
func addCommands(commands ...*AppliedCommand) {
	for _, command := range commands {
		if command.Specification.ID == "" {
			panic("Can not add a command which was not registered")
		}
		registeredCommands[command.Specification.ID] = command
	}
}

// dispatchCommand dispatches a command and return AppliedCommand and ok=true if dispatched successfully
// commandId : id of a command which needs to be dispatched
func dispatchCommand(commandId string) (command *AppliedCommand, ok bool) {
	command, ok = registeredCommands[commandId]

	return
}
