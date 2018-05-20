package dynamic

// ContactListener - обтаботчик контактов
type ContactListener interface {
	BeginContact(с *Contact)
	EndContact(c *Contact)
	PreSolve(c *Contact)
}
