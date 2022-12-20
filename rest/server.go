package main

import (
	"fmt"
	"log"
	"net/http"

	service "github.com/block-listener/services"
)

func main() {
	http.HandleFunc("/getUserInfo", service.GetUserInfo)
	http.HandleFunc("/getTx", service.GetTxwithSignature)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
