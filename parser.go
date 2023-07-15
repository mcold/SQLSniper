package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Snippets struct {
	XMLName xml.Name `xml:"snippets"`
	Groups  []Group  `xml:"group"`
}

type Group struct {
	XMLName  xml.Name  `xml:"group"`
	Category string    `xml:"category,attr"`
	Language string    `xml:"language,attr"`
	Snips    []Snippet `xml:"snippet"`
}

type Snippet struct {
	XMLName     xml.Name `xml:"snippet"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,attr"`
	Code        string   `xml:"code"`
}

func get_snippets(fileURI string) Snippets {
	xmlFile, err := os.Open(fileURI)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var snippets Snippets

	xml.Unmarshal(byteValue, &snippets)

	return snippets
}
