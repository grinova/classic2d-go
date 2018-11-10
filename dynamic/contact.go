package dynamic

import (
	"github.com/grinova/classic2d-go/collision"
	"github.com/grinova/classic2d-go/physics"
)

// ContactFlags - флаги состояния контакта
type ContactFlags int

const (
	// TouchingFlag - соприкосновение происходит
	TouchingFlag ContactFlags = 1 << (iota + 1)
	// WasTouchingFlag - соприкосновение произошло
	WasTouchingFlag
)

// Contact - контакт
type Contact struct {
	BodyA *physics.Body
	BodyB *physics.Body
	Flags ContactFlags
}

// Update обновляет контакт в зависимости от его состояния
func (c *Contact) Update(listener ContactListener) {
	wasToching := c.Flags&TouchingFlag != 0
	if wasToching {
		c.Flags |= WasTouchingFlag
	}
	touching := collision.TestOverlap(c.BodyA, c.BodyB)
	if touching {
		c.Flags |= TouchingFlag
	} else {
		c.Flags &= ^TouchingFlag
	}
	if listener == nil {
		return
	}
	if !wasToching && touching {
		listener.BeginContact(c)
	}
	sensor := false
	if !sensor && touching {
		listener.PreSolve(c)
	}
	if wasToching && !touching {
		listener.EndContact(c)
	}
}
