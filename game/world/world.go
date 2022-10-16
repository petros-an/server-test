package world

import (
	"github.com/petros-an/server-test/common/collider"
	"github.com/petros-an/server-test/common/collider/shape"
	transform "github.com/petros-an/server-test/common/tansform"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	"github.com/petros-an/server-test/game/gameObject"
)

type WorldBorderObject struct {
	gameObject.GameObjectBasic

	transform transform.Transform2D
	Collider  collider.Collider2D
}

func (this *WorldBorderObject) Update(dt float64) {

}

func (this *WorldBorderObject) GetType() gameObject.GameObjectType {
	return gameObject.WorldBorder
}

func (this *WorldBorderObject) GetTransform() transform.Transform2D {
	return this.transform
}

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

func NewWorldBorder(psr vector.PSR2D) *WorldBorderObject {
	newBorder := WorldBorderObject{}
	newBorder.transform.PSR2D = psr

	coll := collider.NewBasic(&newBorder, &shape.HalfPlane{})
	coll.OnCollideField = newBorder.OnCollide
	newBorder.Collider = coll

	return &newBorder
}

func SetUpWorldBorders() []*WorldBorderObject {
	return []*WorldBorderObject{
		NewWorldBorder(vector.PSR2D{Position: vector.New(WorldBorders.Xmax, 0), Scale: vector.New(1, 1), Rotation: vector.New(0, 1)}),
		NewWorldBorder(vector.PSR2D{Position: vector.New(0, WorldBorders.Ymax), Scale: vector.New(1, 1), Rotation: vector.New(-1, 0)}),
		NewWorldBorder(vector.PSR2D{Position: vector.New(WorldBorders.Xmin, 0), Scale: vector.New(1, 1), Rotation: vector.New(0, -1)}),
		NewWorldBorder(vector.PSR2D{Position: vector.New(0, WorldBorders.Ymin), Scale: vector.New(1, 1), Rotation: vector.New(1, 0)}),
	}
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

func (this *WorldBorderObject) OnCollide(gObject gameObject.GameObject) {
	switch gObject.GetType() {
	case gameObject.WorldBorder:
		return

	case gameObject.Character:
		c := gObject.(*character.Character)
		wolrdColl := this.Collider
		halfPlane := wolrdColl.GetShape().(*shape.HalfPlane)

		collisionPoint := halfPlane.InfiniteLine.ClosestPointTo(c.Position())
		c.SetPosition(collisionPoint.Add(halfPlane.Direction.Rotate90().Mul(c.RigidBody.Scale.X/2 + 0.00001)))

	default:
		gObject.Destroy()
	}
}
