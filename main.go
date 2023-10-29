package main

import (
	"fmt"
	"github.com/Sushil-cmd-r/go-microservices/webapp"
	"log"
	"net/http"
)

func main() {
	app := webapp.NewApp()

	productRouter := webapp.NewRouter()

	getProduct := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Get All Products")
		if err != nil {
			log.Fatal("Unable to write to client")
		}
	}

	app.Use("/products", productRouter)

	productRouter.Get("/", getProduct)
	productRouter.Get("/:productID", getProduct)

	app.Listen(":8080")

}
