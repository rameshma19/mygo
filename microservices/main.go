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
