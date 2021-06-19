package main

import (
	"fmt"
	"log"
	"net/http"
)


func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "POshel nahoy")
}
func main() {
	log.Fatal(http.ListenAndServeTLS(":8080","certs/server.crt","certs/server.key", nil))

}
