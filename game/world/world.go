package world

import (
	"github.com/petros-an/server-test/common/vector"
)

type Borders struct {
	Xmin float64
	Xmax float64
	Ymin float64
	Ymax float64
}

var WorldBorders Borders = Borders{
	Xmin: -80,
	Xmax: 80,
	Ymin: -80,
	Ymax: 80,
}

func RestrictPositionWithinBorder(position vector.Vector2D, halfSize vector.Vector2D) vector.Vector2D {
	newPos := position
	if position.X < WorldBorders.Xmin+halfSize.X {
		newPos.X = WorldBorders.Xmin + halfSize.X
	}
	if position.X > WorldBorders.Xmax-halfSize.X {
		newPos.X = WorldBorders.Xmax - halfSize.X
	}
	if position.Y < WorldBorders.Ymin+halfSize.Y {
		newPos.Y = WorldBorders.Ymin + halfSize.Y
	}
	if position.Y > WorldBorders.Ymax-halfSize.Y {
		newPos.Y = WorldBorders.Ymax - halfSize.Y
	}
	// newPos.X = utils.Max(pos.X, BorderXmin)
	// newPos.X = utils.Min(pos.X, BorderXmax)
	// newPos.Y = utils.Max(pos.Y, BorderYmin)
	// newPos.Y = utils.Min(pos.Y, BorderYmax)

	return newPos
}

func IsOutsideWorld(pos vector.Vector2D) bool {
	if pos.X < WorldBorders.Xmin {
		return true
	}
	if pos.X > WorldBorders.Xmax {
		return true
	}
	if pos.Y < WorldBorders.Ymin {
		return true
	}
	if pos.Y > WorldBorders.Ymax {
		return true
	}
	return false
}
