package tavern

import (
	"github.com/Narcolepsick1d/tavern/domain/product"
	"github.com/Narcolepsick1d/tavern/services/order"
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
func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products))
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	uid, err := os.AddCustomer("repcu")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}

}
