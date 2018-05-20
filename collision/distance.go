package collision

import (
	"github.com/grinova/classic2d-go/physics"
)

// Distance - расстояние между телами
func Distance(bodyA *physics.Body, bodyB *physics.Body) float64 {
	r := bodyA.GetRadius() + bodyB.GetRadius()
	d := bodyA.GetPosition().Sub(bodyB.GetPosition()).Length()
	if bodyA.GetInverse() || bodyB.GetInverse() {
		return r*r - d
	}
	return d - r*r
}
