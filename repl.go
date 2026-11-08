package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*GameState, ...string) error
}

func StartGame(gs *GameState) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nCommand >")
		scanner.Scan()
		inputString := scanner.Text()
		cleanedInput := CleanInput(inputString)
		commandString := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}
		if command, ok := getCommands()[commandString]; !ok {
			fmt.Print("Unknown command")
			continue
		} else {
			err := command.callback(gs, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}
}

func CleanInput(text string) []string {
	actualString := strings.TrimPrefix(text, "Command >")
	truncatedString := strings.ToLower(strings.Trim(actualString, " "))
	fields := strings.Fields(truncatedString)
	return fields
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"move": {
			name:     "move",
			callback: commandMove,
		},
	}
}
