package order

import (
	"context"
	"github.com/Narcolepsick1d/tavern/domain/customer"
	"github.com/Narcolepsick1d/tavern/domain/product"

	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Baltika 9", 99.9)
	if err != nil {
		t.Fatal(err)
	}
	nuts, err := product.NewProduct("nuts", "Deez nuts", 69.9)
	if err != nil {
		t.Fatal(err)
	}
	chips, err := product.NewProduct("Layz", "50/50 oxigen and potato", 169.9)
	if err != nil {
		t.Fatal(err)
	}
	return []product.Product{
		beer, nuts, chips,
	}
}
func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("perry")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetId(), order)
	if err != nil {
		t.Error(err)
	}
}
