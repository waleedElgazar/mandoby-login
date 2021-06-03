package main

import (
	"demo/functions"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("")
	http.HandleFunc("/login", functions.Signin)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
