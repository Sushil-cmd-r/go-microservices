package main

import (
	"context"
	"github.com/Sushil-cmd-r/go-microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	LOG := log.New(os.Stdout, "product-api ", log.LstdFlags)
	productHandler := handlers.NewProduct(LOG)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		LOG.Println("Server started on Port 8080\n")
		err := server.ListenAndServe()
		if err != nil {
			LOG.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	LOG.Println("Received terminate, Shutting down gracefully...", sig)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(timeoutContext)
	cancel()

}
