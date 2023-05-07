package order

import (
	"context"
	"github.com/Narcolepsick1d/tavern/domain/customer"
	"github.com/Narcolepsick1d/tavern/domain/customer/memory"
	"github.com/Narcolepsick1d/tavern/domain/customer/mongo"
	"github.com/Narcolepsick1d/tavern/domain/product"
	prodmem "github.com/Narcolepsick1d/tavern/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}
func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}
func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}

}
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}
func (o *OrderService) CreateOrder(customerId uuid.UUID, productsIds []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerId)
	if err != nil {
		return 0, err
	}

	var products []product.Product
	var total float64
	for _, id := range productsIds {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer :%s has order %d products", c.GetId(), len(products))
	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}
	return c.GetId(), nil

}
