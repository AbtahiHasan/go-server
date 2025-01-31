package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w,"hello")
}

func formHandler (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"name = %s\n", name)
	fmt.Fprintf(w,"address = %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("server is listening on port 8080")
}