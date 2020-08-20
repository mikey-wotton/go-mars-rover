package rover

import "fmt"

//Direction describes the way a Rover is facing, using the four cardinal compass points.
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

//Valid will return an error if the direction is not one of the four cardinal directions.
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
