package sim

import (
	//"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	DirPos DirectedPosition
}

func (rob *Robot) Move(room *Room, commands chan rune, actions chan DirectedPosition) {
	cmd := <-commands
	if cmd == Advance {
		rob.DirPos.Dir = North
		if rob.DirPos.Y != room.Length {
			rob.DirPos.Y += 1
		}
	} else if cmd == Left {
		rob.DirPos.Dir = West
		if rob.DirPos.X != 0 {
			rob.DirPos.X -= 1
		}
	} else if cmd == Right {
		rob.DirPos.Dir = East
		if rob.DirPos.X != room.Width {
			rob.DirPos.X += 1
		}
	} else {
		rob.DirPos.Dir = South
		if rob.DirPos.Y != 0 {
			rob.DirPos.Y -= 1
		}
	}
	actions <- rob.DirPos
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}
