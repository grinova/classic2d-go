package vmath

// Vec2 - двухкомпонентный вектор
type Vec2 struct {
	X float64
	Y float64
}

// Add складывает два вектора
func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{X: a.X + b.X, Y: a.Y + b.Y}
}

// Sub вычитает вектор b из вектора a
func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{X: a.X - b.X, Y: a.Y - b.Y}
}

// Mul покомпонентное умножение вектора
func (a Vec2) Mul(b float64) Vec2 {
	return Vec2{X: a.X * b, Y: a.Y * b}
}

// Length возвращает длину вектора
func (a Vec2) Length() float64 {
	return a.X*a.X + a.Y*a.Y
}

// Rotate поворачивает вектор
func (a Vec2) Rotate(rot Rot) Vec2 {
	return Vec2{X: a.X*rot.C - a.Y*rot.S, Y: a.X*rot.S + a.Y*rot.C}
}
