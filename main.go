package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServeTLS(":8080","certs/server.crt","certs/server.key", nil))

}
