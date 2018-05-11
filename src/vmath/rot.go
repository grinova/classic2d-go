package vmath

import (
	"math"
)

// Rot - поворот
type Rot struct {
	C float64
	S float64
}

// RotFromAngle - возвращает вращение для угла в радианах
func RotFromAngle(angle float64) Rot {
	return Rot{C: math.Cos(angle), S: math.Sin(angle)}
}

// RotFromXY - возвращает вращение по координатам
func RotFromXY(x float64, y float64) Rot {
	return Rot{C: x, S: y}
}

// GetAngle возвращает угол в радианах
func (rot Rot) GetAngle() float64 {
	return math.Atan2(rot.S, rot.C)
}

// Inverse возвращает вращение противоположное по знаку
func (rot Rot) Inverse() Rot {
	return Rot{C: rot.C, S: -rot.S}
}

// Normalize возвращает нормализованный поворот
func (rot Rot) Normalize() Rot {
	rl := 1.0 / math.Sqrt(rot.C*rot.C+rot.S*rot.S)
	return Rot{C: rot.C * rl, S: rot.S * rl}
}
