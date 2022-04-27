package discordslash

var (
	// commands is a list of registered SlashedCommand
	registeredCommands = make(map[string]*SlashedCommand)
)

// addCommands adds commands to the dispatcher list
// to dispatch and unregister it later
// commands: command list to register
func addCommands(commands ...*SlashedCommand) {
	for _, command := range commands {
		registeredCommands[command.Name()] = command
	}
}

// dispatchCommand dispatches a command and return SlashedCommand and ok=true if dispatched successfully
// c	: command name
func dispatchCommand(c string) (command *SlashedCommand, ok bool) {
	command, ok = registeredCommands[c]

	return
}

// commandExist checks if command is already added to the list
func commandExist(command string) bool {
	_, ok := dispatchCommand(command)

	return ok
}
