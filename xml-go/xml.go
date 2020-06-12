package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
	Xml string `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

func main() {
	post := Post{
		Id: "1",
		Content: "Hello World",
		Author: Author{
			Id: "2",
			Name: "Sau Sheong",
		},
	}

	output, err := xml.MarshalIndent(&post,"","\t")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("post.xml", output, 0644)
	if err != nil {
		panic(err)
	}
}