// Package classification Products API.
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
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: <none>
//     Contact: Ramesh Mantripragada
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:model
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mygo/microservices/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := NewHello(l)
	// bh := NewGoodBye(l)
	ph := handlers.NewProducts(l)

	server := gin.Default()

	server.GET("/", ph.GetProducts)
	server.POST("/", ph.AddProduct)
	server.PUT("/:id", ph.UpdateProduct)

	server.Run(":9090")
}
