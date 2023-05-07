package customer

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer wasn't found in the repo")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type Repository interface {
	Get(uuid uuid.UUID) (Customer, error)
	Add(customer Customer) error
	Update(customer Customer) error
}
