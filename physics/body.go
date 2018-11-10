package physics

import (
	"github.com/grinova/classic2d-go/physics/shapes"
	"github.com/grinova/classic2d-go/vmath"
)

// BodyType - тип тела
type BodyType int

const (
	// DynamicBody - динамическое тело (перемещаемое)
	DynamicBody BodyType = iota
	// StaticBody - статическое тело (неподвижно)
	StaticBody
)

// Body - тело
type Body struct {
	Type            BodyType
	LinearVelocity  vmath.Vec2
	AngularVelocity float64
	Force           vmath.Vec2
	Torque          float64
	Sweep           vmath.Sweep
	UserData        interface{}
	massData        shapes.MassData
	fixture         Fixture
	xf              vmath.Transform
	radius          float64
	inverse         bool
}

// CreateBody создаёт тело по описанию
func CreateBody(def BodyDef) Body {
	return Body{
		Sweep:           vmath.Sweep{C: def.Position, A: def.Angle},
		xf:              vmath.CreateTransform(def.Position, def.Angle),
		LinearVelocity:  def.LinearVelocity,
		AngularVelocity: def.AngularVelocity,
		inverse:         def.Inverse}
}

// ApplyForce применяет вектор силы
func (b *Body) ApplyForce(force vmath.Vec2) {
	b.Force = b.Force.Add(force)
}

// GetAngle возвращает угол поворота
func (b *Body) GetAngle() float64 {
	return b.xf.Rot.GetAngle()
}

// GetInverse возвращает глаг инвертированности заполнения
func (b *Body) GetInverse() bool {
	return b.inverse
}

// GetMassData возвращает данные массы
func (b *Body) GetMassData() shapes.MassData {
	return b.massData
}

// GetRadius возвращает радиус
func (b *Body) GetRadius() float64 {
	return b.radius
}

// GetPosition возвращает позицию
func (b *Body) GetPosition() vmath.Vec2 {
	return b.xf.Pos
}

// GetRot возвращает вращение
func (b *Body) GetRot() vmath.Rot {
	return b.xf.Rot
}

// SetFixture создаёт и устанавливает фикстуру
func (b *Body) SetFixture(def FixtureDef) Fixture {
	b.fixture = MakeFixture(def)
	if b.fixture.GetDensity() > 0 {
		b.resetMassData()
	}
	b.resetRadius()
	return b.fixture
}

// SetTorque устанавливает крутящий момент
func (b *Body) SetTorque(torque float64) {
	b.Torque = torque
}

// Synchronize - синхронизация трансформации
func (b *Body) Synchronize() {
	b.xf = vmath.CreateTransform(b.Sweep.C, b.Sweep.A)
}

func (b *Body) resetMassData() {
	b.massData = b.fixture.GetMassData()
	b.massData.Center = b.massData.Center.Mul(1.0 / b.massData.Mass)
}

func (b *Body) resetRadius() {
	radius := b.fixture.GetShape().GetRadius()
	if b.inverse {
		b.radius = -radius
	} else {
		b.radius = radius
	}
}
