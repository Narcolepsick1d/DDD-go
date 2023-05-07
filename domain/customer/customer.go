package customer

import (
	"errors"
	"github.com/Narcolepsick1d/tavern"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have valid name")
)

type Customer struct {
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &tavern.Person{
		Name: name,
		Id:   uuid.New(),
	}
	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}
func (c *Customer) GetId() uuid.UUID {
	return c.person.Id
}
func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Id = id
}
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}
func (c *Customer) GetName() string {
	return c.person.Name
}
