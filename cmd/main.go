package main

import (
	"context"
	"github.com/Narcolepsick1d/tavern/domain/product"
	"github.com/Narcolepsick1d/tavern/services/order"
	"github.com/Narcolepsick1d/tavern/services/tavern"
	"github.com/google/uuid"
	"log"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products))
	if err != nil {
		log.Fatal(err)
	}
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		log.Fatal(err)
	}
	uid, err := os.AddCustomer("Percy")
	if err != nil {
		log.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}
func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Baltika 9", 99.9)
	if err != nil {
		log.Fatal(err)
	}
	nuts, err := product.NewProduct("nuts", "Deez nuts", 69.9)
	if err != nil {
		log.Fatal(err)
	}
	chips, err := product.NewProduct("Layz", "50/50 oxigen and potato", 169.9)
	if err != nil {
		log.Fatal(err)
	}
	return []product.Product{
		beer, nuts, chips,
	}
}
