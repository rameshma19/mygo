// Package handlers Products API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
// Schemes: http
// Host: localhost
// BasePath: /
// Version: 0.0.1
// License: <none>
// Contact: Ramesh Mantripragada
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
//
// swagger:model
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mygo/microservices/data"
)

// Products struct
type Products struct {
	l *log.Logger
}

// NewProducts
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products listProducts
func (p *Products) GetProducts(c *gin.Context) {
	p.l.Println("Handle Get Products using Gin framework")
	lp := data.GetProducts()
	rw := c.Writer
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route POST /products AddProducts
func (p *Products) AddProduct(ctx *gin.Context) {
	p.l.Println("Handle POST Products using Gin framework")
	rw := ctx.Writer

	prod := &data.Product{}
	err := prod.FromJSON(ctx.Request.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	ctx.BindJSON(prod)
	data.AddProduct(prod)

	p.l.Printf("Prod: %#v", prod)
}

// swagger:route PUT /products UpdateProducts
func (p *Products) UpdateProduct(c *gin.Context) {

	p.l.Println("Handle PUT Products using Gin framework")

	rw := c.Writer
	req := c.Request

	prod := &data.Product{}

	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	idS, _ := c.Params.Get("id")
	p.l.Println("id", idS)
	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(rw, "Unable to convert id to int", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Prod: %#v", prod)
}

//Just trying git
