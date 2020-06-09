package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	Content string	`json:"content"`
	Author Author	`json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `jsomn:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great time",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "Have a good day",
				Author:  "Betty",
			},
		},
	}

	jsonFile, _ := os.Create("post3.json")
	encoder := json.NewEncoder(jsonFile)
	encoder.Encode(&post)
	fmt.Println(post)
}