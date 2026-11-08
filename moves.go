package main

import (
	"fmt"
	"strings"
)

func commandMove(gs *GameState, args ...string) error {
	// example "move a2 to b3"
	err := validateMoveCommand(args)
	if err != nil {
		return err
	}
	oldX, oldY, newX, NewY := getCoordinates(args)
	fmt.Printf("%d, %d, %d, %d", oldX, oldY, newX, NewY)
	return nil
}

func getCoordinates(words []string) (oldX, oldY, newX, newY int) {
	xAxisLookupTable := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
	}

	yAxisLookupTable := map[string]int{
		"1": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
	}
	oldX = xAxisLookupTable[strings.Split(words[0], "")[0]]
	oldY = yAxisLookupTable[strings.Split(words[0], "")[1]]
	newX = xAxisLookupTable[strings.Split(words[1], "")[0]]
	newY = yAxisLookupTable[strings.Split(words[1], "")[1]]

	return
}

func validateMoveCommand(words []string) error {
	for _, s := range words {
		if len(s) != 2 || s[0] < 'a' || s[0] > 'h' || s[1] < '1' || s[1] > '8' {
			return fmt.Errorf("wrong format of move command")
		}
	}
	return nil
}
