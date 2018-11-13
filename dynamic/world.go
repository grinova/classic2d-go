package dynamic

import (
	"math"
	"time"

	"github.com/grinova/classic2d-server/physics"
	"github.com/grinova/classic2d-server/vmath"
)

type worldFlags int

const (
	newBodies = 1 << (iota + 1)
	clearForces
	defaultFlags = newBodies
)

// World - мир
type World struct {
	bodies         *bodies
	contactManager ContactManager
	flags          worldFlags
}

// CreateWorld создаёт мир
func CreateWorld() World {
	world := World{bodies: &bodies{}, flags: defaultFlags}
	world.contactManager = CreateContactManager(&world)
	return world
}

// ClearForces обнуляет значение действующей силы для всех тел
func (w *World) ClearForces() {
	for list := w.bodies.first; list != nil; list = list.Next {
		list.Body.Force = vmath.Vec2{}
	}
}

// CreateBody создаёт тело по описанию
func (w *World) CreateBody(def physics.BodyDef) *physics.Body {
	w.flags |= newBodies
	body := physics.CreateBody(def)
	w.bodies.add(&body)
	return &body
}

// Clear очищает мир
func (w *World) Clear() {
	for list := w.bodies.first; list != nil; list = list.Next {
		w.DestroyBody(list.Body)
	}
	w.contactManager.Clear()
	w.flags = defaultFlags
}

// DestroyBody утичтожает тело
func (w *World) DestroyBody(body *physics.Body) {
	w.bodies.remove(body)
}

// GetBodies возвращает список всех тел
func (w *World) GetBodies() *BodyList {
	return w.bodies.first
}

// GetContactManager возвращает менеджер контактов
func (w *World) GetContactManager() *ContactManager {
	return &w.contactManager
}

// SetContactListener устанавливает обработчик контактов
func (w *World) SetContactListener(listener ContactListener) {
	w.contactManager.SetContactListener(listener)
}

// Step - шаг расчета процессов за время time миллисекунд
func (w *World) Step(d time.Duration) {
	ms := d.Seconds() * 1000
	if ms > 100 {
		return
	}
	if w.flags&newBodies != 0 {
		w.contactManager.FindNewContacts()
		w.flags &= ^newBodies
	}

	iterations := math.Floor(ms / math.Min(ms, 4.0))
	T := ms / (iterations * 1000)
	for i := 0.0; i < iterations; i++ {
		for list := w.bodies.first; list != nil; list = list.Next {
			body := list.Body
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

		for list := w.bodies.first; list != nil; list = list.Next {
			list.Body.Synchronize()
		}
	}

	if w.flags&clearForces != 0 {
		w.ClearForces()
	}
}
