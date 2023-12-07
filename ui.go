package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func (app *config) makeUI() (*widget.List, *widget.List, *widget.Entry, *widget.Entry) {

	app.Snips = get_snippets("UserSnippets.xml")

	arr := []string{}
	snips := []string{}

	app.retrim()
	for gi := range app.Snips.Groups {
		arr = append(arr, app.Snips.Groups[gi].Category)
	}

	for si := range app.Snips.Groups[0].Snips {
		snips = append(snips, app.Snips.Groups[0].Snips[si].Name)
	}

	app.SnipsArr = snips
	app.ListGrData = binding.BindStringList(&arr)

	l_groups := widget.NewListWithData(app.ListGrData,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	app.ListSnData = binding.BindStringList(&app.SnipsArr)

	l_snips := widget.NewListWithData(app.ListSnData,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// Snippet
	app.EditSnip = app.Snips.Groups[0].Snips[0].Code
	app.EditSnipData = binding.BindString(&app.EditSnip)

	edit := widget.NewMultiLineEntry()
	edit.Bind(app.EditSnipData)

	// Description
	app.EditDescr = app.Snips.Groups[0].Snips[0].Description
	app.EditDescrData = binding.BindString(&app.EditDescr)

	editDescr := widget.NewMultiLineEntry()
	editDescr.Bind(app.EditDescrData)

	l_groups.OnSelected = app.refreshGroup
	l_snips.OnSelected = app.refreshSnip

	app.IDGroup = 0
	app.IDSnip = 0

	l_groups.Select(app.IDGroup)
	l_snips.Select(app.IDSnip)

	return l_groups, l_snips, edit, editDescr
}
