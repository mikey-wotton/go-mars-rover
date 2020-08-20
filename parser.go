package parser

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mikey-wotton/go-mars-rover/rover"
	"strconv"
	"strings"
)

var (
	ErrEmptyInput               = errors.New("input is empty")
	ErrRoverWithoutInstructions = errors.New("rover missing instructions")
	ErrBoundariesNotProvided    = errors.New("two boundaries are required")
	ErrInvalidBoundary          = errors.New("invalid boundary provided")
	ErrRoverInitialise          = errors.New("rover initialise not provided x, y, and direction")
)

const (
	numBoundaries      = 2 //X, Y
	numRoverInitValues = 3 //X, Y, and Direction
)

func ParseInstructions(input string) (rover.Rovers, error) {
	return nil, nil
}

func parseBoundary(scanner *bufio.Scanner) (*rover.Coordinate, error) {
	line := scanner.Text()

	strs := strings.Split(line, " ")
	if len(strs) != numBoundaries {
		return nil, ErrBoundariesNotProvided
	}

	boundX, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil, ErrInvalidBoundary
	}

	boundY, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, ErrInvalidBoundary
	}

	return &rover.Coordinate{
		X: boundX,
		Y: boundY,
	}, nil
}

func parseRoverPosition(scanner *bufio.Scanner) (*rover.Position, error) {
	line := scanner.Text()

	strs := strings.Split(line, " ")
	if len(strs) != numRoverInitValues {
		return nil, ErrRoverInitialise
	}

	posX, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil, fmt.Errorf("x boundary not supplied : %w", err)
	}

	posY, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, fmt.Errorf("y boundary not supplied : %w", err)
	}

	dir, err := stringToDirection(strs[2])
	if err != nil {
		return nil, err
	}

	return &rover.Position{
		Coordinate: rover.Coordinate{
			X: posX,
			Y: posY,
		},
		Direction: dir,
	}, nil
}

func stringToDirection(s string) (rover.Direction, error) {
	var dir rover.Direction

	switch s {
	case "North":
		dir = rover.North
	case "East":
		dir = rover.East
	case "South":
		dir = rover.South
	case "West":
		dir = rover.West
	default:
		return rover.UnknownDirection, fmt.Errorf("unknown direction string %s", s)
	}

	return dir, nil
}