package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var xmlStart string

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

	// sorting groups
	sort.Slice(snippets.Groups, func(i, j int) bool {
		return strings.ToLower(snippets.Groups[i].Category) < strings.ToLower(snippets.Groups[j].Category)
	})

	// sorting snippets
	for gi := range snippets.Groups {
		sort.Slice(snippets.Groups[gi].Snips, func(i, j int) bool {
			return snippets.Groups[gi].Snips[i].Name < snippets.Groups[gi].Snips[j].Name
		})
	}

	return snippets
}

func (app *config) retrim() {
	// trim leading / tailing spaces
	for gi := range app.Snips.Groups {
		for si := range app.Snips.Groups[gi].Snips {
			app.Snips.Groups[gi].Snips[si].Code = strings.TrimSpace(app.Snips.Groups[gi].Snips[si].Code)
		}
	}
}

func replace_content(fileURI string) {
	xmlStart = "<?xml version = '1.0' encoding = 'UTF-8'?>"

	contents, _ := ioutil.ReadFile(fileURI)
	contents = []byte(strings.Replace(string(contents), "&#xA;", "\n", 100500))
	contents = []byte(strings.Replace(string(contents), "&#39;", "'", 100500))
	contents = []byte(strings.Replace(string(contents), "<code>", "<code><![CDATA[\n", 100500))
	contents = []byte(strings.Replace(string(contents), "</code>", "]]>\n</code>", 100500))

	contents = []byte(xmlStart + "\n" + string(contents))

	ioutil.WriteFile(fileURI, contents, 0644)
}
