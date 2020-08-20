package parser

import (
	"bufio"
	"errors"
	"github.com/mikey-wotton/go-mars-rover/rover"
)

var (
	ErrEmptyInput               = errors.New("input is empty")
	ErrRoverWithoutInstructions = errors.New("rover missing instructions")
	ErrBoundariesNotProvided    = errors.New("two boundaries are required")
	ErrInvalidBoundary          = errors.New("invalid boundary provided")
	ErrRoverInitialise          = errors.New("rover initialise not provided x, y, and direction")
)


func ParseInstructions(input string) (rover.Rovers, error) {
	return nil, nil
}

func parseBoundary(scanner *bufio.Scanner) (*rover.Coordinate, error) {
	return nil, nil
}

func parseRoverPosition(scanner *bufio.Scanner) (*rover.Position, error) {
	return nil, nil
}