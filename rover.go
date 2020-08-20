package rover

import (
	"errors"
)

var (
	ErrBoundaryNorth          = errors.New("rover at Y edge cannot move north")
	ErrBoundaryEast           = errors.New("rover at X edge cannot move east")
	ErrBoundarySouth          = errors.New("rover at Y edge cannot move south")
	ErrBoundaryWest           = errors.New("rover at X edge cannot move west")
	ErrRoverOutsideXBoundary  = errors.New("rover x coordinate must be within boundary")
	ErrRoverOutsideYBoundary  = errors.New("rover x coordinate must be within boundary")
	ErrRoverRequiresCommands  = errors.New("rover must have at least one valid command")
	ErrRoverNotInitialised    = errors.New("rover must not be nil")
	ErrPositionNotInitialised = errors.New("rover position must not be nil")
)

type Rovers []*Rover

type Position struct {
	Coordinate
	Direction
}

type Rover struct {
	Commands string
	Position *Position
	Boundary *Coordinate
}

func (r *Rover) Explore() error {
	return nil
}

func (r *Rover) Valid() error {

	return nil
}

func (r *Rover) move() error {
	return nil
}

func (r *Rover) turn(i Instruction) error {
	return nil
}
