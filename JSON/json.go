package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		Id: 1,
		Content: "Hello World",
		Author: Author{
			Id: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id: 3,
				Content: "Have a great time",
				Author: "Adam",
			},
			{
				Id: 4,
				Content: "Have a good day",
				Author: "Betty",
			},
		},
	}

	output, _ := json.MarshalIndent(&post, "", "\t")
	fmt.Println(string(output))
	//Marshalは構造体をもとにjson形式に変換してくれる
	err := ioutil.WriteFile("post2.json", output, 0644)
	//marshalされた変数outputをWriteFileの引数に渡してpost2.jsonに書き込む
	if err != nil {
		panic(err)
	}
}