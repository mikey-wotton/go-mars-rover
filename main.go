package main

import (
	"fmt"
	"github.com/mikey-wotton/go-mars-rover/parser"
	"log"
)

const input = `5 5
1 2 North
LMLMLMLMM
3 3 East
MMRMMRMRRM`

func main() {
	rovers, err := parser.ParseInstructions(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rovers {
		fmt.Println(fmt.Sprintf("Starting Position (%d, %d) Facing %s", r.Position.X, r.Position.Y, r.Position.Direction.String()))
		fmt.Println(fmt.Sprintf("Instructions: %s", r.Commands))
		err := r.Explore()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Finishing Position (%d, %d) Facing %s", r.Position.X, r.Position.Y, r.Position.Direction.String()))
		fmt.Println()
	}
}
