package shapes

// ShapeType - тип формы
type ShapeType int

const (
	// Circle - круг
	Circle ShapeType = iota
	// Polygon - полигон
	Polygon
)

// Shape - форма
type Shape interface {
	ComputeMass(density float64) MassData
	GetRadius() float64
	GetType() ShapeType
}
