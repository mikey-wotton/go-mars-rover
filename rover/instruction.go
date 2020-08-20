package rover

import "fmt"

//Instruction represents the available movements a Rover can perform.
type Instruction int32

const (
	Move      Instruction = 'M'
	TurnLeft  Instruction = 'L'
	TurnRight Instruction = 'R'
)

//Valid will return an error if the current Instruction is not one of the three available Rover Instructions.
func (i Instruction) Valid() error {
	switch i {
	case Move:
	case TurnLeft:
	case TurnRight:
	default:
		return fmt.Errorf("rover provided unknown Instruction{%d}", i)
	}

	return nil
}
