package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error Occurred", http.StatusBadRequest)
			//responseWriter.WriteHeader(http.StatusBadRequest)
			//val, _ := responseWriter.Write([]byte("Error Occurred"))
			//fmt.Printf("Val: %d", val)
			return
		}
		_, _ = fmt.Fprintf(responseWriter, "Hello %s\n", body)
	})

	http.HandleFunc("/goodbye", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Println("Bye World")
	})

	_ = http.ListenAndServe("localhost:8080", nil)
}
