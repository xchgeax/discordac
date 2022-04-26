package discordslash

var (
	// commands is a list of registered SlashedCommand
	commands = make(map[string]SlashedCommand)
)

// addCommand adds a command to the dispatcher list commands
// to dispatch and unregister it later
func addCommand(command SlashedCommand) {
	commands[command.Name()] = command
}

// dispatchCommand dispatches a command and return SlashedCommand and ok=true if dispatched successfully
// c	: command name
func dispatchCommand(c string) (command SlashedCommand, ok bool) {
	command, ok = commands[c]

	return
}

// commandExist checks if command is already added to the list
func commandExist(command string) bool {
	_, ok := dispatchCommand(command)

	return ok
}
