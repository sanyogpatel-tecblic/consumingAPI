package consuming

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Consuming() {
	url := "http://localhost:8010/tasks"
	//NewRequest is used to create a new HTTP request with the method, URL, and optional request body.

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//http - client is used to send the response and retrive the response
	client := &http.Client{}
	//here we are passing the req to the client so that we can send and recieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//ioutil.ReadAll is used to read the response body of an HTTP request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
