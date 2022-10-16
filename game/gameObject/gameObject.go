package gameObject

import transform "github.com/petros-an/server-test/common/tansform"

type GameObjectType int

const (
	Character = iota
	Projectile
	WorldBorder
)

type GameObject interface {
	Update(dt float64)
	GetType() GameObjectType
	ToDestroy() bool
	Destroy()
	GetTransform() transform.Transform2D
}

type GameObjectBasic struct {
	toDestroy bool
}

func (this GameObjectBasic) ToDestroy() bool {
	return this.toDestroy
}

func (this *GameObjectBasic) Destroy() {
	this.toDestroy = true
}
