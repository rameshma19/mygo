package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mygo/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(rw, req)
		return
	}

	if req.Method == http.MethodPut {
		//extract the ID in the URI
		p.l.Println("In PUT")
		path := req.URL.Path
		rx := regexp.MustCompile(`/([0-9]+)`)
		exp := rx.FindAllStringSubmatch(path, -1)

		if len(exp) != 1 {
			http.Error(rw, "Invalid URI1", http.StatusBadRequest)
			return
		}

		if len(exp[0]) != 2 {
			http.Error(rw, "Invalid URI2", http.StatusBadRequest)
			return
		}

		idString := exp[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI3", http.StatusBadRequest)
		}
		p.l.Println("got id", id)

		p.updateProduct(id, rw, req)
		return

	}
	// all other case
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Get Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Post Products")

	prod := &data.Product{}

	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Post Products")

	prod := &data.Product{}

	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Prod: %#v", prod)
}

//Just trying git
