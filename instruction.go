package rover

import "fmt"

type Instruction int32

const (
	Move      Instruction = 'M'
	TurnLeft  Instruction = 'L'
	TurnRight Instruction = 'R'
)

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
