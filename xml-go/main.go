package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
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

type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil{
		panic(err)
	}
	defer xmlFile.Close()
	//xmlData, err := ioutil.ReadAll(xmlFile)
	//if err != nil {
	//	panic(err)
	//}
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF{
			break
		}
		if err != nil {
			panic(err)
		}

		switch se := t.(type) {
		case xml.StartElement:

			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}
	////fmt.Println(string(xmlData))
	//var post Post
	//xml.Unmarshal(xmlData, &post)
	//fmt.Println(post)
}
