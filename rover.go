package rover

import (
	"errors"
	"fmt"
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
	if r == nil {
		return ErrRoverNotInitialised
	}

	//check boundaries
	switch {
	case r.Boundary.X < 0:
		return fmt.Errorf("rover has a negative x boundary %d but should not", r.Boundary.X)
	case r.Boundary.Y < 0:
		return fmt.Errorf("rover has a negative y boundary %d but should not", r.Boundary.Y)
	}

	//check positions
	switch {
	case r.Position == nil:
		return ErrPositionNotInitialised
	case r.Position.X < 0:
		return ErrRoverOutsideXBoundary
	case r.Position.X > r.Boundary.X:
		return ErrRoverOutsideXBoundary
	case r.Position.Y < 0:
		return ErrRoverOutsideYBoundary
	case r.Position.Y > r.Boundary.Y:
		return ErrRoverOutsideYBoundary
	}

	//check direction
	if err := r.Position.Direction.Valid(); err != nil {
		return err
	}

	//check instructions
	if len(r.Commands) < 1 {
		return ErrRoverRequiresCommands
	}
	for _, command := range r.Commands {
		if err := Instruction(command).Valid(); err != nil {
			return err
		}
	}

	return nil
}

func (r *Rover) move() error {
	return nil
}

func (r *Rover) turn(i Instruction) error {
	return nil
}
