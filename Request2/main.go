package main

import (
	"fmt"
	"net/http"
)

func  header(w http.ResponseWriter, r *http.Request){
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	bytes := make([]byte, len)
	r.Body.Read(bytes)
	fmt.Println(string(bytes))
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", header)
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}