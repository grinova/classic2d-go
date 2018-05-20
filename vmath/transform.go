package vmath

// Transform - трансформация
type Transform struct {
	Pos Vec2
	Rot Rot
}

// CreateTransform создаёт трансформацию по позиции и углу в радианах
func CreateTransform(pos Vec2, angle float64) Transform {
	return Transform{Pos: pos, Rot: RotFromAngle(angle)}
}
