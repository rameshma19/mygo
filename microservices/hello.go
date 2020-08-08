package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("Hello World!!")
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", d)
}
