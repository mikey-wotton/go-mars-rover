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

//Position details the location and direction a Rover is on Mars.
type Position struct {
	Coordinate
	Direction
}


//Rover represents a rover which is used to explore the Mars surface.
type Rover struct {
	Commands string
	Position *Position
	Boundary *Coordinate
}

//Explore is used to execute the instructions that belong to the rover, allowing it to traverse the Mars surface
//up to its boundaries, if the Rover cannot perform an instruction it will return an error.
func (r *Rover) Explore() error {
	for _, command := range r.Commands {
		instruction := Instruction(command)
		switch instruction {
		case Move:
			if err := r.move(); err != nil {
				return err
			}
		case TurnLeft:
			if err := r.turn(instruction); err != nil {
				return err
			}
		case TurnRight:
			if err := r.turn(instruction); err != nil {
				return err
			}
		default:
			return fmt.Errorf("rover provided unknown Instruction{%d}", instruction)
		}
	}

	return nil
}

//Valid will return an error if the Rover is in a non-valid state, such as out of boundaries or facing an unknown
//direction.
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
	switch r.Position.Direction {
	case North:
		if r.Position.Y+1 <= r.Boundary.Y {
			r.Position.Y += 1
		} else {
			return ErrBoundaryNorth
		}
	case East:
		if r.Position.X+1 <= r.Boundary.X {
			r.Position.X += 1
		} else {
			return ErrBoundaryEast
		}
	case South:
		if r.Position.Y-1 >= 0 {
			r.Position.Y -= 1
		} else {
			return ErrBoundarySouth
		}
	case West:
		if r.Position.X-1 >= 0 {
			r.Position.X -= 1
		} else {
			return ErrBoundaryWest
		}
	default:
		return errUnknownDirection(r.Position.Direction)
	}

	return nil
}

func (r *Rover) turn(i Instruction) error {
	switch i {
	case TurnLeft:
		newPos := r.Position.Direction - 1
		if newPos >= 1 {
			r.Position.Direction = newPos
		} else {
			r.Position.Direction = West
		}
	case TurnRight:
		r.Position.Direction += 1
		if r.Position.Direction > 4 {
			r.Position.Direction = North
		}
	default:
		return fmt.Errorf("unknown instruction passed to update direction %v", i)
	}

	return nil
}
