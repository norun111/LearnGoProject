package main

import (
	"encoding/xml"
	"os"
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

func main() {
	post := Post{
		Id: "1",
		Content: "Hello World",
		Author: Author{
			Id: "2",
			Name: "Sau Sheong",
		},
	}

	xmlFile, err := os.Create("post2.xml")
	if err != nil {
		panic(err)
	}
	encoder := xml.NewEncoder(xmlFile)
	err = encoder.Encode(&post)
	if err != nil {
		panic(err)
	}
}
