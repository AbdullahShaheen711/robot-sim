package sim

import (
	"fmt"
	"math/rand"
	"time"
)

type Room struct {
	Width  int
	Length int
}

func (room *Room) IssueReport(actions chan DirectedPosition) {
	action := <-actions
	fmt.Printf("At {%v,%v} Heading %c\n", action.X, action.Y, action.Dir)
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}
