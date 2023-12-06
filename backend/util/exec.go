package util

import (
	"errors"
	"os/exec"
)

/**
 * Execute a command
 * Simplified version of exec.Command
 */
func Exec(args ...string) ([]byte, error) {
	var command *exec.Cmd
	if len(args) == 0 {
		return []byte{}, errors.New("No command provided")
	} else if len(args) == 1 {
		command = exec.Command(args[0])
	} else {
		command = exec.Command(args[0], args[1:]...)
	}
	return command.CombinedOutput()
}
