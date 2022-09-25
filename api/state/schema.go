package state

type DirectionInput struct {
	X float64
	Y float64
}

type PlayerInput struct {
	Direction *DirectionInput
}
