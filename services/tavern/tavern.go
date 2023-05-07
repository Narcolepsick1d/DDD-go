package tavern

import (
	"github.com/Narcolepsick1d/tavern/services/order"
	"github.com/google/uuid"
	"log"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService *order.OrderService
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}
func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf(" \nbill the customer :%0.0f \n", price)
	return nil

}
