package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body")
		w.WriteHeader(500)
		return
	}
	r.Body.Close()

	fmt.Printf("URL: %s, Remote Address: %s; Method: %s\n", r.URL.Path, r.RemoteAddr, r.Method)
	fmt.Println(string(body))
	w.WriteHeader(200)
}
