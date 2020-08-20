package rover

import "fmt"

type Direction uint8

//go:generate stringer -type=Direction
const (
	UnknownDirection Direction = iota
	North
	East
	South
	West
)



func errUnknownDirection(d Direction) error {
	return fmt.Errorf("rover facing unknown direction %v", d)
}

func (d Direction) Valid() error {
	switch d {
	case North:
	case East:
	case South:
	case West:
	default:
		return errUnknownDirection(d)
	}

	return nil
}
