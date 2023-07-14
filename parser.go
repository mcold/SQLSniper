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

func get_snippets() Snippets {
	xmlFile, err := os.Open("UserSnippets.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var snippets Snippets

	xml.Unmarshal(byteValue, &snippets)

	// for i := 0; i < len(snippets.Groups); i++ {
	// 	fmt.Println("Category: " + snippets.Groups[i].Category)
	// 	for j := 0; j < len(snippets.Groups[i].Snippets); j++ {
	// 		fmt.Println("Snippet name: " + snippets.Groups[i].Snippets[j].Name)
	// 		fmt.Println("Snippets code: " + snippets.Groups[i].Snippets[j].Code)
	// 	}
	// }

	return snippets
}