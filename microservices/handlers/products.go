package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mygo/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

// 	if req.Method == http.MethodGet {
// 		p.getProducts(rw, req)
// 		return
// 	}

// 	if req.Method == http.MethodPost {
// 		p.addProduct(rw, req)
// 		return
// 	}

// 	if req.Method == http.MethodPut {
// 		//extract the ID in the URI
// 		p.l.Println("In PUT")
// 		path := req.URL.Path
// 		rx := regexp.MustCompile(`/([0-9]+)`)
// 		exp := rx.FindAllStringSubmatch(path, -1)

// 		if len(exp) != 1 {
// 			http.Error(rw, "Invalid URI1", http.StatusBadRequest)
// 			return
// 		}

// 		if len(exp[0]) != 2 {
// 			http.Error(rw, "Invalid URI2", http.StatusBadRequest)
// 			return
// 		}

// 		idString := exp[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw, "Invalid URI3", http.StatusBadRequest)
// 		}
// 		p.l.Println("got id", id)

// 		p.updateProduct(id, rw, req)
// 		return

// 	}
// 	// all other case
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *Products) GetProducts(c *gin.Context) {
	p.l.Println("Handle Get Products using Gin framework")
	lp := data.GetProducts()
	rw := c.Writer
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(c *gin.Context) {
	p.l.Println("Handle POST Products using Gin framework")
	rw := c.Writer

	prod := &data.Product{}
	err := prod.FromJSON(c.Request.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) UpdateProduct(c *gin.Context) {

	p.l.Println("Handle PUT Products using Gin framework")

	rw := c.Writer
	req := c.Request

	prod := &data.Product{}

	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(c.Param("id"))
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

func (p *Products) MiddlewareValidateProduct(c *gin.Context) {

	p.l.Println("Inside MiddlewareValidateProduct using Gin framework")

	rw := c.Writer
	req := c.Request
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	return
}

// func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
// 	p.l.Println("Handle Get Products")

// 	lp := data.GetProducts()

// 	err := lp.ToJSON(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// 	}
// }

// func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
// 	p.l.Println("Handle Post Products")

// 	prod := &data.Product{}

// 	err := prod.FromJSON(req.Body)

// 	if err != nil {
// 		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
// 	}

// 	data.AddProduct(prod)

// 	p.l.Printf("Prod: %#v", prod)
// }

// func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {

// 	vars := mux.Vars(req)

// 	p.l.Println("Handle Put1 Products")

// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(rw, "Unable to convert id to int", http.StatusBadRequest)
// 		return
// 	}

// 	p.l.Println("Handle Put2 Products", id)

// 	prod := &data.Product{}

// 	err = prod.FromJSON(req.Body)

// 	if err != nil {
// 		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
// 	}

// 	err = data.UpdateProduct(id, prod)
// 	if err != nil {
// 		http.Error(rw, "Product Not found", http.StatusInternalServerError)
// 		return
// 	}

// 	p.l.Printf("Prod: %#v", prod)
// }

//Just trying git
