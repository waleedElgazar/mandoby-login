package main

import (
	"demo/functions"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("")
	http.HandleFunc("/verify", functions.Signin)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
