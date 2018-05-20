package shapes

import (
	"math"
)

// CircleShape - круг
type CircleShape struct {
	Radius float64 // = -1
}

// ComputeMass - вычисление данных массы
func (s CircleShape) ComputeMass(density float64) MassData {
	mass := density * 2 * math.Pi * s.Radius
	return MassData{Mass: mass}
}

// GetRadius возвращает радиус круга
func (s CircleShape) GetRadius() float64 {
	return s.Radius
}

// GetType возвращает тип фигуры
func (s CircleShape) GetType() ShapeType {
	return Circle
}
