package parser

import (
	"fmt"
	"github.com/mikey-wotton/go-mars-rover/rover"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInstructions(t *testing.T) {
	tests := map[string]struct {
		input     string
		expRovers rover.Rovers
		expErr    error
	}{
		"single rover test": {
			input: `1 1
0 0 South
M`,
			expRovers: rover.Rovers{
				&rover.Rover{
					Boundary: &rover.Coordinate{X: 1, Y: 1},
					Commands: "M",
					Position: &rover.Position{
						Coordinate: rover.Coordinate{},
						Direction:  rover.South,
					},
				},
			},
			expErr: nil,
		},
		"multiple rover test": {
			input: `5 5
1 2 North
LMLMLMLMM
3 3 West
LLLLRRRR`,
			expRovers: rover.Rovers{
				&rover.Rover{
					Boundary: &rover.Coordinate{X: 5, Y: 5},
					Commands: "LMLMLMLMM",
					Position: &rover.Position{
						Coordinate: rover.Coordinate{X: 1, Y: 2},
						Direction:  rover.North,
					},
				},
				&rover.Rover{
					Boundary: &rover.Coordinate{X: 5, Y: 5},
					Commands: "LLLLRRRR",
					Position: &rover.Position{
						Coordinate: rover.Coordinate{X: 3, Y: 3},
						Direction:  rover.West,
					},
				},
			},
			expErr: nil,
		},
		"example rover test": {
			input: `5 5
1 2 North
LMLMLMLMM
3 3 East
MMRMMRMRRM`,
			expRovers: rover.Rovers{
				&rover.Rover{
					Boundary: &rover.Coordinate{X: 5, Y: 5},
					Commands: "LMLMLMLMM",
					Position: &rover.Position{
						Coordinate: rover.Coordinate{X: 1, Y: 2},
						Direction:  rover.North,
					},
				},
				&rover.Rover{
					Boundary: &rover.Coordinate{X: 5, Y: 5},
					Commands: "MMRMMRMRRM",
					Position: &rover.Position{
						Coordinate: rover.Coordinate{X: 3, Y: 3},
						Direction:  rover.East,
					},
				},
			},
			expErr: nil,
		},
		"err rover outside X boundary": {
			input: `1 1
2 1 North
LLLMMMRRR`,
			expRovers: nil,
			expErr:    rover.ErrRoverOutsideXBoundary,
		},
		"err rover outside Y boundary": {
			input: `1 1
1 2 North
LLLMMMRRR`,
			expRovers: nil,
			expErr:    rover.ErrRoverOutsideYBoundary,
		},
		"err rover invalid direction": {
			input: `1 1
1 1 Northeasterly
LLLMMMRRR`,
			expRovers: nil,
			expErr:    fmt.Errorf("unknown direction string %s", "Northeasterly"),
		},
		"err no instructions found": {
			input:     ``,
			expRovers: nil,
			expErr:    ErrEmptyInput,
		},
		"err rover invalid instruction": {
			input: `1 1
0 0 South
LRMX`,
			expRovers: nil,
			expErr:    fmt.Errorf("rover provided unknown Instruction{%d}", 'X'),
		},
		"err rover without instructions": {
			input: `1 1
0 0 South`,
			expRovers: nil,
			expErr:    ErrRoverWithoutInstructions,
		},
	}

	for description, test := range tests {
		rovers, err := ParseInstructions(test.input)
		assert.Equalf(t, test.expErr, err, "%s failed, expected error %v but got %v", description, test.expErr, err)
		assert.ElementsMatchf(t, test.expRovers, rovers, "%s failed, expected rovers %v but got %v", description, test.expRovers, rovers)
	}
}
