package dynamic

// ContactListener - обтаботчик контактов
type ContactListener interface {
	BeginContact(c *Contact)
	EndContact(c *Contact)
	PreSolve(c *Contact)
}
