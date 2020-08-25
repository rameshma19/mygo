package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func FindProduct(id int) (*Product, int, error) {
	for _, p := range productList {
		if p.ID == id {
			return p, 0, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func GetNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
