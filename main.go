package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/domwakeling/go_server/handlers"
	"github.com/domwakeling/go_server/data"
)

func main() {
	newLogger := log.New(os.Stdout, "product-api", log.LstdFlags)
	validator := data.NewValidation()
	productsHandler := handlers.NewProducts(newLogger, validator)

	// use mux to make a new router
	muxRouter := mux.NewRouter()

	// set up individual subrouters for each method and add handlers
	getRouter := muxRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.GetProducts)
	
	putRouter := muxRouter.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := muxRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	deleteRouter := muxRouter.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProduct)

	// handle docs
	opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	specHandler := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", specHandler)
	getRouter.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	// redirect root to /docs
	getRouter.HandleFunc("/", productsHandler.RootHandler)

	server := &http.Server{
		Addr: ":9090",
		Handler: muxRouter,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// wrap the ListenAndServe in a go func so that it doesn't block ...
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			newLogger.Fatal(err)
		}
	}()

	// make a new channel, send it a broadcast if os interrupts or kills
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	// block until there is a broadcast on the sigChan; then log
	sig := <- sigChan
	newLogger.Println("\n***\nReceived terminate, graceful shutdown:", sig)

	// the shutdown won't trigger until a signal is broadcast and the above element unblocks
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(ctx)
}
