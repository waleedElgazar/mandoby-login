package main

import (
	"demo/functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/verify", functions.Signin)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
