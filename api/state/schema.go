package state

type DirectionSchema struct {
	X float64
	Y float64
}

type PlayerInputSchema struct {
	Direction       *DirectionSchema
	ProjectileFired *ProjectileFiredSchema
}

type ProjectileFiredSchema struct {
	Direction *DirectionSchema
}
