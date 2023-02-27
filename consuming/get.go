package consuming

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Task struct {
	ID     int    `json:"id"`
	Tasks  string `json:"tasks"`
	UserID int    `json:"userid"`
}

func MakeGet() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:8010/tasks"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}

		//now send and retrive the response for that we will be using http.Client()

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var tasks []Task

		if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.ParseFiles("home.page.tmpl"))

		err = tmpl.Execute(w, tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8060", nil))
}
