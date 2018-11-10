package dynamic

import (
	"github.com/grinova/classic2d-server/physics"
	"github.com/grinova/classic2d-server/vmath"
)

// ContactSolver - разрешатель контактов
type ContactSolver struct {
	world World
}

// Solve разрешает контакты
func (cs ContactSolver) Solve() {
	contactManager := cs.world.GetContactManager()
	contacts := contactManager.GetContacts()
	for contact := range contacts {
		if contact.Flags&WasTouchingFlag != 0 {
			continue
		}
		bodyA, bodyB := contact.BodyA, contact.BodyB
		switch {
		case bodyA.Type == physics.DynamicBody && bodyB.Type == physics.DynamicBody:
			solveDynamic(bodyA, bodyB)
		case bodyA.Type == physics.DynamicBody && bodyB.Type == physics.StaticBody:
			solveStatic(bodyA, bodyB)
		case bodyA.Type == physics.StaticBody && bodyB.Type == physics.DynamicBody:
			solveStatic(bodyB, bodyA)
		}
	}
}

func solveDynamic(bodyA *physics.Body, bodyB *physics.Body) {
	vA := bodyA.LinearVelocity
	vB := bodyB.LinearVelocity
	massDataA := bodyA.GetMassData()
	massDataB := bodyB.GetMassData()
	mA := massDataA.Mass
	mB := massDataB.Mass

	mcA := massDataA.Center.Add(bodyA.Sweep.C)
	mcB := massDataB.Center.Add(bodyB.Sweep.C)
	x := mcB.X - mcA.X
	y := mcB.Y - mcA.Y
	massRot := vmath.RotFromXY(x, -y).Normalize()

	vA = vA.Rotate(massRot)
	vB = vB.Rotate(massRot)

	uAx := (mB*(2*vB.X-vA.X) + mA*vA.X) / (mA + mB)
	uAy := vA.Y
	uBx := (mA*(2*vA.X-vB.X) + mB*vB.X) / (mA + mB)
	uBy := vB.Y

	uA := vmath.Vec2{X: uAx, Y: uAy}
	uB := vmath.Vec2{X: uBx, Y: uBy}

	massRot = massRot.Inverse()
	uA = uA.Rotate(massRot)
	uB = uB.Rotate(massRot)

	bodyA.LinearVelocity = uA
	bodyB.LinearVelocity = uB
}

func solveStatic(bodyA *physics.Body, bodyB *physics.Body) {
	cA := bodyA.Sweep.C
	cB := bodyB.Sweep.C
	vA := bodyA.LinearVelocity
	massDataA := bodyA.GetMassData()
	massDataB := bodyB.GetMassData()

	mcA := massDataA.Center.Add(cA)
	mcB := massDataB.Center.Add(cB)
	x := mcB.X - mcA.X
	y := mcB.Y - mcA.Y
	massRot := vmath.RotFromXY(x, -y).Normalize()

	vA = vA.Rotate(massRot)

	uAx := -vA.X
	uAy := vA.Y

	uA := vmath.Vec2{X: uAx, Y: uAy}

	massRot = massRot.Inverse()
	uA = uA.Rotate(massRot)

	bodyA.LinearVelocity = uA
}
