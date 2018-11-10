package physics

import (
	"github.com/grinova/classic2d-server/physics/shapes"
)

// FixtureDef - описание фикстуры
type FixtureDef struct {
	Shape   shapes.Shape
	Density float64
}
