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

func RestrictPositionWithinBorder(position vector.Vector2D, halfSize vector.Vector2D) vector.Vector2D {

	newPos := position
	if position.X < BorderXmin+halfSize.X {
		newPos.X = BorderXmin + halfSize.X
	}
	if position.X > BorderXmax-halfSize.X {
		newPos.X = BorderXmax - halfSize.X
	}
	if position.Y < BorderYmin+halfSize.Y {
		newPos.Y = BorderYmin + halfSize.Y
	}
	if position.Y > BorderYmax-halfSize.Y {
		newPos.Y = BorderYmax - halfSize.Y
	}
	// newPos.X = utils.Max(pos.X, BorderXmin)
	// newPos.X = utils.Min(pos.X, BorderXmax)
	// newPos.Y = utils.Max(pos.Y, BorderYmin)
	// newPos.Y = utils.Min(pos.Y, BorderYmax)

	return newPos
}

func IsOutsideWorld(pos vector.Vector2D) bool {
	if pos.X < BorderXmin {
		return true
	}
	if pos.X > BorderXmax {
		return true
	}
	if pos.Y < BorderYmin {
		return true
	}
	if pos.Y > BorderYmax {
		return true
	}
	return false
}
