package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func WriteExample(w http.ResponseWriter, r *http.Request){
	str := `<html>
				<head><title>Go Web Prog</title></head>
				<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

func headerExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User: "tomoya",
		Threads: []string{"first","second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main(){

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}