package collision

import (
	"github.com/grinova/classic2d-server/physics"
	"github.com/grinova/classic2d-server/settings"
)

// TestOverlap проверяет пересечение тел
func TestOverlap(bodyA *physics.Body, bodyB *physics.Body) bool {
	return Distance(bodyA, bodyB) < 10*settings.EPSILON
}
