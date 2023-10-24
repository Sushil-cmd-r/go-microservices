package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	LOG *log.Logger
}

func NewHello(LOG *log.Logger) *Hello {
	return &Hello{LOG: LOG}
}

func (hello *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	hello.LOG.Println("Hello World")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Error Occurred", http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(responseWriter, "Hello %s\n", body)
}
