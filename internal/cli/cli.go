// Package cli holds primitives common to all slk commands
package cli

// Command abstracts slk command
type Command interface {
	// Run runs the command.
	Run()
	// Usage prints command usage
	Usage()
}

type initFunc func() Command

var commands map[string]initFunc

// RegisterCommand writes command init function to commands map
func RegisterCommand(name string, initCommand initFunc) {
	if commands == nil {
		commands = make(map[string]initFunc)
	}
	commands[name] = initCommand
}

// InitCommands returns all registered commands
func InitCommands() map[string]initFunc {
	return commands
}
