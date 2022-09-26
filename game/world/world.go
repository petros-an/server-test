package world

import (
	"github.com/petros-an/server-test/common/vector"
)

const (
	BorderXmin = -40
	BorderXmax = 40
	BorderYmin = -40
	BorderYmax = 40
)

func RestrictPositionWithinBorder(pos vector.Vector2D) vector.Vector2D {

	newPos := pos
	if pos.X < BorderXmin {
		newPos.X = BorderXmin
	}
	if pos.X > BorderXmax {
		newPos.X = BorderXmax
	}
	if pos.Y < BorderYmin {
		newPos.Y = BorderYmin
	}
	if pos.Y > BorderYmax {
		newPos.Y = BorderYmax
	}
	// newPos.X = utils.Max(pos.X, BorderXmin)
	// newPos.X = utils.Min(pos.X, BorderXmax)
	// newPos.Y = utils.Max(pos.Y, BorderYmin)
	// newPos.Y = utils.Min(pos.Y, BorderYmax)

	return newPos
}
