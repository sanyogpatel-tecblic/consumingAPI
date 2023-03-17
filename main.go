package main

import (
	"fmt"
	"net/http"

	"github.com/sanyogpatel-tecblic/consuming-API/consuming"
)

func main() {
	// consuming.Consuming()
	consuming.MakeGet()
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	fmt.Println("Listening on port 8020...")
}
