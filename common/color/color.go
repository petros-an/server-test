package color

import "math/rand"

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

func Random() RGBColor {
	return RGBColor{
		R: uint8(rand.Intn(155) + 100),
		B: uint8(rand.Intn(155) + 100),
		G: uint8(rand.Intn(155) + 100),
	}
}
