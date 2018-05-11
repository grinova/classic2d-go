package physics

import (
	"vmath"
)

// BodyDef - орписание тела
type BodyDef struct {
	Position        vmath.Vec2
	Angle           float64
	LinearVelocity  vmath.Vec2
	AngularVelocity float64
	Inverse         bool
}
