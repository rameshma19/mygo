package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price" validate:"gt=0"`
	SKU       string  `json:"sku"`
	CreatedAt string  `json:"~"`
	UpdatedAt string  `json:"~"`
	DeletedAt string  `json:"~"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)

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
