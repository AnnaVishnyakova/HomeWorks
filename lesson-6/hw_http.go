package main

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
func HelloWorldByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Println(w, "Hello, %s\n", name)
}
func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/name", HelloWorldByNameHandler)
	http.ListenAndServe(":8082", nil)
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", fs)
	// http.ListenAndServe(":80", nil)
}
