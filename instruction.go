package rover

type Instruction int32

const (
	Move      Instruction = 'M'
	TurnLeft  Instruction = 'L'
	TurnRight Instruction = 'R'
)

func (i Instruction) Valid() error {
	return nil
}
