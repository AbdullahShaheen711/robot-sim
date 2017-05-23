package sim

import ()

type DirectedPosition struct {
	X, Y int
	Dir  rune
}

const (
	North rune = 'N'
	South rune = 'S'
	East  rune = 'E'
	West  rune = 'W'
)

const (
	Left    rune = 'L'
	Advance rune = 'A'
	Right   rune = 'R'
	Back    rune = 'B'
)
