package main

import (
	"devsimiyu/azurefunc/common"
	"devsimiyu/azurefunc/hellofunc"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.TODO()

	mongo := common.NewMongoConnection(ctx)

	defer func() {
		if err := mongo.Disconnect(ctx); err != nil {
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello-v1", hellofunc.HelloHandler)

	var address string

	port, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !ok {
		port = "8080"
	}

	address = ":" + port

	log.Println("HelloFunc server listening on " + address)
	log.Fatal(http.ListenAndServe(address, mux))
}
