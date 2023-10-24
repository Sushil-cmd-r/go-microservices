package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct {
	LOG *log.Logger
}

func NewGoodbye(logger *log.Logger) *GoodBye {
	return &GoodBye{
		LOG: logger,
	}
}

func (goodbye *GoodBye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	goodbye.LOG.Println("Goodbye world")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Error Occurred", http.StatusBadRequest)
		return
	}

	_, _ = fmt.Fprintf(responseWriter, "Bye %s\n", body)

}
