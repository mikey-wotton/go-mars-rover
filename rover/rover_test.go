package rover

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRover_Explore(t *testing.T) {
	tests := map[string]struct {
		rover       *Rover
		expErr      error
		expPosition *Position
	}{
		"example first rover explores": {
			rover:&Rover{
					Boundary: &Coordinate{X: 5, Y: 5},
					Commands: "LMLMLMLMM",
					Position: &Position{
						Coordinate: Coordinate{X: 1, Y: 2},
						Direction:  North,
					},
				},
			expErr: nil,
			expPosition: &Position{
				Coordinate: Coordinate{1,3},
				Direction:  North,
			},
		},
		"example second rover explores": {
			rover:&Rover{
				Boundary: &Coordinate{X: 5, Y: 5},
				Commands: "MMRMMRMRRM",
				Position: &Position{
					Coordinate: Coordinate{X: 3, Y: 3},
					Direction:  East,
				},
			},
			expErr: nil,
			expPosition: &Position{
				Coordinate: Coordinate{5,1},
				Direction:  East,
			},
		},
		"explore to boundary edge diagonally south to north": {
			rover: &Rover{
				Commands: "RMLMRMLMR",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: North,
				},
				Boundary: &Coordinate{
					X: 2,
					Y: 2,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 2,
					Y: 2,
				},
				Direction: East,
			},
			expErr: nil,
		},
		"explore entire boundary starting North/West going south": {
			rover: &Rover{
				Commands: "MMMMMLMMMMMLMMMMMLMMMMM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 5,
					},
					Direction: South,
				},
				Boundary: &Coordinate{
					X: 5,
					Y: 5,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 5,
				},
				Direction: West,
			},
			expErr: nil,
		},
		"turn right multiple times without moving": {
			rover: &Rover{
				Commands: "RRRRR",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: North,
				},
				Boundary: &Coordinate{
					X: 2,
					Y: 2,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 0,
				},
				Direction: East,
			},
			expErr: nil,
		},
		"turn left multiple times without moving": {
			rover: &Rover{
				Commands: "LLLLL",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: South,
				},
				Boundary: &Coordinate{
					X: 2,
					Y: 2,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 0,
				},
				Direction: East,
			},
			expErr: nil,
		},
		"err trying to leave northern boundary": {
			rover: &Rover{
				Commands: "MM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: North,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 1,
				},
				Direction: North,
			},
			expErr: ErrBoundaryNorth,
		},
		"err trying to leave eastern boundary": {
			rover: &Rover{
				Commands: "MM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: East,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 1,
					Y: 0,
				},
				Direction: East,
			},
			expErr: ErrBoundaryEast,
		},
		"err trying to leave southern boundary": {
			rover: &Rover{
				Commands: "MM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: South,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 0,
				},
				Direction: South,
			},
			expErr: ErrBoundarySouth,
		},
		"err trying to leave western boundary": {
			rover: &Rover{
				Commands: "MM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: West,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 0,
				},
				Direction: West,
			},
			expErr: ErrBoundaryWest,
		},
		"err dealing with unknown instruction": {
			rover: &Rover{
				Commands: "MX",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: East,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 1,
					Y: 0,
				},
				Direction: East,
			},
			expErr: fmt.Errorf("rover provided unknown Instruction{%d}", 'X'),
		},
		"err rover facing unknown direction, does not move": {
			rover: &Rover{
				Commands: "MM",
				Position: &Position{
					Coordinate: Coordinate{
						X: 0,
						Y: 0,
					},
					Direction: UnknownDirection,
				},
				Boundary: &Coordinate{
					X: 1,
					Y: 1,
				},
			},
			expPosition: &Position{
				Coordinate: Coordinate{
					X: 0,
					Y: 0,
				},
				Direction: UnknownDirection,
			},
			expErr: errUnknownDirection(UnknownDirection),
		},
	}

	for desc, test := range tests {
		err := test.rover.Explore()
		assert.Equalf(t, test.expErr, err, "%s failed, expected %v but got %v", desc, test.expErr, err)
		assert.Equalf(t, test.expPosition.X, test.rover.Position.X, "%s failed, expected x position %d but got %d", desc, test.expPosition.X, test.rover.Position.X)
		assert.Equalf(t, test.expPosition.Y, test.rover.Position.Y, "%s failed, expected y position %d but got %d", desc, test.expPosition.Y, test.rover.Position.Y)
		assert.Equalf(t, test.expPosition.Direction, test.rover.Position.Direction, "%s failed, expected direction %s but got %s", desc, string(test.expPosition.Direction), string(test.rover.Position.Direction))
	}
}

func TestRover_Valid(t *testing.T) {
	tests := map[string]struct {
		rover  *Rover
		expErr error
	}{
		"nil error if rover is in valid state": {
			rover:  &Rover{
				Commands: "LMR",
				Position: &Position{
					Coordinate: Coordinate{1,1},
					Direction:  South,
				},
				Boundary: &Coordinate{2,2},
			},
			expErr: nil,
		},
		"err if rover not init": {
			rover:  nil,
			expErr: ErrRoverNotInitialised,
		},
		"err if rover.Boundary.X negative value throws error": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{1, 1},
					Direction:  255,
				},
				Boundary: &Coordinate{-1, 1},
			},
			expErr: fmt.Errorf("rover has a negative x boundary %d but should not", -1),
		},
		"err if rover.Boundary.Y negative value throws error": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{1, 1},
					Direction:  255,
				},
				Boundary: &Coordinate{1, -1},
			},
			expErr: fmt.Errorf("rover has a negative y boundary %d but should not", -1),
		},
		"err if rover.Position not init": {
			rover: &Rover{
				Commands: "LLL",
				Position: nil,
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrPositionNotInitialised,
		},
		"err if rover.Position.X less than 0": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{-1, 0},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrRoverOutsideXBoundary,
		},
		"err if rover.Position.X greater than boundary.X": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{2, 0},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrRoverOutsideXBoundary,
		},
		"err if rover.Position.Y less than 0": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{0, -1},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrRoverOutsideYBoundary,
		},
		"err if rover.Position.Y greater than boundary.Y": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{1, 2},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrRoverOutsideYBoundary,
		},
		"err if rover.Position.Direction is not valid": {
			rover: &Rover{
				Commands: "LLL",
				Position: &Position{
					Coordinate: Coordinate{1, 1},
					Direction:  255,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: errUnknownDirection(Direction(255)),
		},
		"err if rover.Commands has no commands": {
			rover: &Rover{
				Commands: "",
				Position: &Position{
					Coordinate: Coordinate{1, 1},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: ErrRoverRequiresCommands,
		},
		"err if rover.Commands has invalid command": {
			rover: &Rover{
				Commands: "LLLX",
				Position: &Position{
					Coordinate: Coordinate{1, 1},
					Direction:  North,
				},
				Boundary: &Coordinate{1, 1},
			},
			expErr: fmt.Errorf("rover provided unknown Instruction{%d}", Instruction('X')),
		},
	}

	for desc, test := range tests {
		err := test.rover.Valid()
		assert.Equalf(t, test.expErr, err, "%s failed, expected %v but got %v", desc, test.expErr, err)
	}
}
