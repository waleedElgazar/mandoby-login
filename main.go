package main

import (
	"demo/functions"
	"log"
	"net/http"
)

func main() {
	port := functions.GetPort()
	http.HandleFunc("/", functions.Welcome)
	http.HandleFunc("/verify", functions.Signin)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
