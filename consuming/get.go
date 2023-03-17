package consuming

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageurl"`
}

func MakeGet() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Make a GET request to the API endpoint
		resp, err := http.Get("http://localhost:8020/items")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Unmarshal the JSON response into a struct
		var Items []Item
		err = json.NewDecoder(resp.Body).Decode(&Items)
		if err != nil {
			log.Fatal(err)
		}

		// Render an HTML page with the data
		tmpl := template.Must(template.ParseFiles("./templates/home.page.tmpl"))
		err = tmpl.Execute(w, Items)
		if err != nil {
			log.Fatal(err)
		}
	})
	// Serve static files
	log.Fatal(http.ListenAndServe(":4080", nil))
}
