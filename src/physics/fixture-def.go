package physics

import (
	"physics/shapes"
)

// FixtureDef - описание фикстуры
type FixtureDef struct {
	Shape   shapes.Shape
	Density float64
}
