package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	// Run the Server
	runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API")
	if runtime_api != "" {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(router)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		log.Println("Starting up on own")
		srv := &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: router,
		}
		_ = srv.ListenAndServe()
	}
}
