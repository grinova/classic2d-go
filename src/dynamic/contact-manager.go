package dynamic

import (
	"collision"
	"physics"
)

// Contacts - контакты
type Contacts map[*Contact]bool

// ContactManager - менеджер контактов
type ContactManager struct {
	world    *World
	contacts Contacts
	listener ContactListener
}

// CreateContactManager создаёт менеджер контактов
func CreateContactManager(world *World) ContactManager {
	return ContactManager{world: world, contacts: make(Contacts)}
}

// AddPair добавляет контакт для пары тел
func (cm *ContactManager) AddPair(bodyA *physics.Body, bodyB *physics.Body) *Contact {
	contact := &Contact{BodyA: bodyA, BodyB: bodyB}
	cm.contacts[contact] = true
	return contact
}

// Collide разрешает столкновения тел
func (cm *ContactManager) Collide() {
	for contact := range cm.contacts {
		contact.Update(cm.listener)
		overlap := collision.TestOverlap(contact.BodyA, contact.BodyB)
		if !overlap {
			cm.Destroy(contact)
		}
	}
}

// Clear удаляет все контакты
func (cm *ContactManager) Clear() {
	for contact := range cm.contacts {
		cm.Destroy(contact)
	}
}

// Destroy удаляет контакт
func (cm *ContactManager) Destroy(contact *Contact) {
	delete(cm.contacts, contact)
}

// FindNewContacts - поиск новых контактов
func (cm *ContactManager) FindNewContacts() {
	bodies := cm.world.GetBodies()
	for bodyA := range bodies {
		for bodyB := range bodies {
			if bodyA == bodyB {
				continue
			}
			if collision.TestOverlap(bodyA, bodyB) && !cm.hasContact(bodyA, bodyB) {
				cm.AddPair(bodyA, bodyB)
			}
		}
	}
}

// GetContacts возвращает контакты
func (cm *ContactManager) GetContacts() Contacts {
	return cm.contacts
}

// SetContactListener устанавливает обработчик контактов
func (cm *ContactManager) SetContactListener(listener ContactListener) {
	cm.listener = listener
}

func (cm *ContactManager) hasContact(bodyA *physics.Body, bodyB *physics.Body) bool {
	for contact := range cm.contacts {
		if contact.BodyA == bodyA && contact.BodyB == bodyB ||
			contact.BodyB == bodyA && contact.BodyA == bodyB {
			return true
		}
	}
	return false
}
