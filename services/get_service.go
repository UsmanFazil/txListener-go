package service

import (
	"fmt"
	"net/http"
)

func GetUserTx(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id =>", id)

	fmt.Fprintf(w, "Hello!")
}
