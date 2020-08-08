package main

import (
	"log"
	"net/http"
)

type Products2 struct {
	l *log.Logger
}

func newProductHandler(l *log.Logger) *Products2 {
	return &Products2{l}
}

func (p *Products2) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	} else if req.Method == http.MethodPost {
		p.addProduct(rw, req)
		return
	}
	// all other case
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products2) getProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Get Products")

	lp := getProductList()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products2) addProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Post Products")
}

//Just trying git
