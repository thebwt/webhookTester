package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	SaveData(r)
	fmt.Fprint(w, "Message Acknowledged")
	//Grab variables: headers, multipart form payload, URI
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{.*}", HomeHandler)

	http.ListenAndServe("0.0.0.0:8080", r)
}

//Takes an incoming request and stores it somehow
// in this case I'll be using mgo
func SaveData(r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("This is working! %s\n", r.RequestURI)
	fmt.Printf( "Second line! %s\n", body)
	fmt.Printf( "Third line! %s", r.Header)

}
