package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mygo/microservices/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	//env.Parse()

	// http.HandleFunc("/hello", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Hello World!!")
	// })

	// http.HandleFunc("/goodbye", func(rw http.ResponseWriter, req *http.Request) {

	// 	// log.Printf("Data %s\n", d)
	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := NewHello(l)
	// bh := NewGoodBye(l)
	ph := handlers.NewProducts(l)

	//sm := mux.NewRouter()
	// getRouter := sm.Methods(http.MethodGet).Subrouter()
	// getRouter.HandleFunc("/", ph.GetProducts)

	server := gin.Default()

	server.GET("/", ph.GetProducts)
	server.POST("/", ph.AddProduct)
	server.PUT("/:id", ph.UpdateProduct)

	server.Run(":9090")

	// putRouter := sm.Methods(http.MethodPut).Subrouter()
	// putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	// postRouter := sm.Methods(http.MethodPost).Subrouter()
	// postRouter.HandleFunc("/", ph.AddProduct)

	//sm := http.NewServeMux()
	// sm.Handle("/hello", hh)
	// sm.Handle("/bye", bh)
	//sm.Handle("/products", ph)

	// MyHTTPSvr := &http.Server{
	// 	Addr:         ":8081",
	// 	Handler:      sm,
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// }
	// go func() {
	// 	err := MyHTTPSvr.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(l)
	// 	}
	// }()

	// sigChan := make(chan os.Signal)
	// signal.Notify(sigChan, os.Interrupt)
	// signal.Notify(sigChan, os.Kill)

	// sig := <-sigChan
	// l.Println("Received Terminate, graceful shutdown", sig)

	// tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// //MyHTTPSvr.Shutdown(tc)
}

// func reqHandler(http.ResponseWriter, *http.Request) {
// 	log.Println("Hello World!!")
// }
