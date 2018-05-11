package collision

import (
	"physics"
	"settings"
)

// TestOverlap проверяет пересечение тел
func TestOverlap(bodyA *physics.Body, bodyB *physics.Body) bool {
	return Distance(bodyA, bodyB) < 10*settings.EPSILON
}
