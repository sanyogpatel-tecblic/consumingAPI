package main

import (
	"net/http"

	"github.com/sanyogpatel-tecblic/consuming-API/consuming"
)

func main() {
	// consuming.Consuming()
	http.HandleFunc("/", consuming.Handler)
	http.Handle("/uploads/", http.StripPrefix("./uploads/", http.FileServer(http.Dir("uploads"))))
	http.ListenAndServe(":8050", nil)
}
