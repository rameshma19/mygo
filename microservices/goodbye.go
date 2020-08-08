package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("Goodbye World!!")
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))
		return
	}
	fmt.Fprintf(rw, "Goodbye %s\n", d)
}
