package main

import (
	"fmt"
	"sim"
)

func main() {
	// initialize commands channel between Controller and Robot
	cmdChan := make(chan rune)
	// initialize actions channel between Robot and Room
	actChan := make(chan sim.DirectedPosition)
	// initialize controller with commands list, robot and room
	var controller *sim.Controller = &sim.Controller{[]rune{'A', 'L', 'A', 'A', 'R', 'R', 'B', 'A', 'L', 'A', 'L'}}
	// robot starts at{0,0} facing north
	var robot *sim.Robot = &sim.Robot{sim.DirectedPosition{0, 0, sim.North}}
	// room is 5 X 6
	var room *sim.Room = &sim.Room{5, 6}
	// main target
	go controller.Command(cmdChan)
	go robot.Move(room, cmdChan, actChan)
	room.IssueReport(actChan)
	fmt.Println("Main is over!")
}
