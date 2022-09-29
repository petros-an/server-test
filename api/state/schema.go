package state

type PlayerInputSchema struct {
	MoveDirection     *MoveDirectionSchema
	ProjectileInput   *ProjectileFiredSchema
	CharacterRotation *CharacterRotationSchema
}

type MoveDirectionSchema struct {
	X float64
	Y float64
}

type CharacterRotationSchema struct {
	X float64
	Y float64
}

type ProjectileFiredSchema struct {
	X float64
	Y float64
}
