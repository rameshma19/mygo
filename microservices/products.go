package main

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        int     `jason: "id"`
	Name      string  `jason: "name"`
	Desc      string  `jason: "desc"`
	Price     float32 `jason: "price"`
	SKU       string  `jason: "sku"`
	CreatedAt string  `jason: "~"`
	UpdatedAt string  `jason: "~"`
	DeletedAt string  `jason: "~"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func getProductList() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "Coffee with milk",
		Price:     2.50,
		SKU:       "IMN101",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Coffee without milk",
		Price:     3.50,
		SKU:       "IMN102",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}
