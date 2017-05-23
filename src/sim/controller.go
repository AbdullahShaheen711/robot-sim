package sim

import (
	//"fmt"
	"math/rand"
	"time"
)

type Controller struct {
	Commands []rune
}

func (controller *Controller) Command(cmdChan chan rune, cmd rune) {
	cmdChan <- cmd
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}
