package sim

import (
	"fmt"
)

type Controller struct {
	Commands []rune
}

func (controller *Controller) Command(cmdChan chan rune) {
	for _, cmd := range controller.Commands {
		cmdChan <- cmd
	}
	close(cmdChan)
	fmt.Println("Commands channel closes")
}
