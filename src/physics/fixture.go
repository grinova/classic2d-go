package physics

import (
	"physics/shapes"
)

// Fixture - фикстура
type Fixture struct {
	shape   shapes.Shape
	density float64
}

// MakeFixture создаёт фикстуру по описанию
func MakeFixture(def FixtureDef) Fixture {
	return Fixture{shape: def.Shape, density: def.Density}
}

// GetDensity возвращает протность
func (f Fixture) GetDensity() float64 {
	return f.density
}

// GetShape возвращает форму
func (f Fixture) GetShape() shapes.Shape {
	return f.shape
}

// GetMassData возвращает данные массы
func (f Fixture) GetMassData() shapes.MassData {
	return f.shape.ComputeMass(f.density)
}
