package collision

import (
	"github.com/grinova/classic2d-go/physics"
	"github.com/grinova/classic2d-go/settings"
)

// TestOverlap проверяет пересечение тел
func TestOverlap(bodyA *physics.Body, bodyB *physics.Body) bool {
	return Distance(bodyA, bodyB) < 10*settings.EPSILON
}
