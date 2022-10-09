package gameObject

import transform "github.com/petros-an/server-test/common/tansform"

type GameObject interface {
	Update(dt float64)
	GetType() int
	ToDestroy() bool
	Destroy()
	GetTransform() transform.Transform2D
}

const (
	Character = iota
	Projectile
)
