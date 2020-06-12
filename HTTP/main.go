package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func process(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form["username"])

	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err==nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go</title></head>
<body><p>Hello World</p></body>
</html>`

	w.Write([]byte(str))
}

type Post struct {
	User string
	Threads []string
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, " there not")
}

func headerExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User: "tomoyaueno",
		Threads: []string{"1","2","3"},
	}
	json, _ := json.Marshal(&post)
	w.Write(json)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeExample", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}