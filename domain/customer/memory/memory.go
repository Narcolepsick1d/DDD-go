package memory

import (
	"fmt"
	"github.com/Narcolepsick1d/tavern/domain/customer"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}
func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}
func (mr *MemoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Lock()
	}
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exist: %w", customer.ErrFailedToAddCustomer)

	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}
func (mr *MemoryRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}
