package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "This is working! %s\n", r.RequestURI)
	fmt.Fprintf(w, "Second line! %s\n", body)
	fmt.Fprintf(w, "Third line! %s", r.Header)
	//Grab variables: headers, multipart form payload, URI
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{.*}", HomeHandler)

	http.ListenAndServe("0.0.0.0:8080", r)
}
