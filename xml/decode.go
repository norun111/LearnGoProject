package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main(){

	xmlFile, err := os.Open("post.xml")
	if err != nil{
		return
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)

	for {
		t, err := decoder.Token()
	}

}
