package dynamic

import (
	"math"
	"physics"
	"vmath"
)

type worldFlags int

const (
	newBodies = 1 << (iota + 1)
	clearForces
	defaultFlags = newBodies
)

// Bodies - тела
type Bodies map[*physics.Body]bool

// World - мир
type World struct {
	bodies         Bodies
	contactManager ContactManager
	flags          worldFlags
}

// CreateWorld создаёт мир
func CreateWorld() World {
	world := World{bodies: make(Bodies), flags: defaultFlags}
	world.contactManager = CreateContactManager(&world)
	return world
}

// ClearForces обнуляет значение действующей силы для всех тел
func (w *World) ClearForces() {
	for body := range w.bodies {
		body.Force = vmath.Vec2{}
	}
}

// CreateBody создаёт тело по описанию
func (w *World) CreateBody(def physics.BodyDef) *physics.Body {
	w.flags |= newBodies
	body := physics.CreateBody(def)
	w.bodies[&body] = true
	return &body
}

// Clear очищает мир
func (w *World) Clear() {
	for body := range w.bodies {
		w.DestroyBody(body)
	}
	w.contactManager.Clear()
	w.flags = defaultFlags
}

// DestroyBody утичтожает тело
func (w *World) DestroyBody(body *physics.Body) {
	delete(w.bodies, body)
}

// GetBodies возвращает список всех тел
func (w *World) GetBodies() Bodies {
	return w.bodies
}

// GetContactManager возвращает менеджер контактов
func (w *World) GetContactManager() ContactManager {
	return w.contactManager
}

// SetContactListener устанавливает обработчик контактов
func (w *World) SetContactListener(listener ContactListener) {
	w.contactManager.SetContactListener(listener)
}

// Step - шаг расчета процессов за время time миллисекунд
func (w *World) Step(time float64) {
	if time > 100 {
		return
	}
	if w.flags&newBodies != 0 {
		w.contactManager.FindNewContacts()
		w.flags &= ^newBodies
	}

	iterations := math.Floor(time / math.Min(time, 4.0))
	T := time / (iterations * 1000)
	for i := 0.0; i < iterations; i++ {
		for body := range w.bodies {
			if body.Type == physics.StaticBody {
				continue
			}
			a := body.Force.Mul(T)
			vs := body.LinearVelocity.Mul(T)
			as := a.Mul(T * T / 2)
			pos := body.GetPosition().Add(vs).Add(as)
			body.AngularVelocity += body.Torque * T
			da := body.AngularVelocity * T

			body.LinearVelocity = body.LinearVelocity.Add(a)
			body.Sweep.C = pos
			body.Sweep.A = body.GetAngle() + da
		}

		w.contactManager.FindNewContacts()
		w.contactManager.Collide()
		contactSolver := ContactSolver{world: *w}
		contactSolver.Solve()

		for body := range w.bodies {
			body.Synchronize()
		}
	}

	if w.flags&clearForces != 0 {
		w.ClearForces()
	}
}
