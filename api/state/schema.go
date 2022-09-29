package state

type PlayerInputSchema struct {
	MoveDirection   *DirectionSchema
	ProjectileInput *ProjectileFiredSchema
}

type DirectionSchema struct {
	X float64
	Y float64
}

type ProjectileFiredSchema struct {
	X float64
	Y float64
}
