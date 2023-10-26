package handlers

import (
	"errors"
	"github.com/Sushil-cmd-r/go-microservices/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	LOG *log.Logger
}

func NewProduct(logger *log.Logger) *Product {
	return &Product{
		LOG: logger,
	}
}

func (product *Product) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		product.GetProducts(responseWriter, request)
		return
	}
	if request.Method == http.MethodPost {
		product.AddProduct(responseWriter, request)
		return
	}
	if request.Method == http.MethodPut {
		regx := regexp.MustCompile(`/([0-9]+)`)
		path := request.URL.Path
		group := regx.FindAllStringSubmatch(path, -1)

		if len(group) != 1 {
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := group[0][1]
		id, _ := strconv.Atoi(idString)

		product.LOG.Println("Got ID: ", id)

		product.updateProduct(id, responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Product) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	products := data.GetProducts()
	responseWriter.Header().Add("Content-Type", "application/json")

	//d, err := json.Marshal(products)
	//_, _ = responseWriter.Write(d)

	err := products.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Error: Unable to marshal data", http.StatusInternalServerError)
		return
	}
}

func (product *Product) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	//responseWriter.WriteHeader(http.StatusNotImplemented)
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (product *Product) updateProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	} else {
		http.Error(responseWriter, "something went wrong", http.StatusInternalServerError)
		return
	}
}
