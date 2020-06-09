package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id int `json:"id"`
	Content string	`json:"content"`
	Author Author	`json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id string `jsomn:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()
	jsonData, _ := ioutil.ReadAll(jsonFile)

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}