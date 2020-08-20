package rover

import "fmt"

type Direction uint8

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
	return nil
}
