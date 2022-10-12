package gameObject

import transform "github.com/petros-an/server-test/common/tansform"

type GameObjectType string

const (
	Character  = "character"
	Projectile = "projectile"
)

type GameObject interface {
	Update(dt float64)
	GetType() GameObjectType
	ToDestroy() bool
	Destroy()
	GetTransform() transform.Transform2D
}
