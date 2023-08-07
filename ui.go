package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (app *config) makeUI() (*widget.Entry, *widget.Button, *widget.Entry, *widget.List, *widget.List) {

	app.Snips = get_snippets("UserSnippets.xml")

	app.retrim()

	filter := widget.NewEntry()
	btn := widget.NewButton("Find", func() { app.searchFilter(app.FilterWidget.Text) })
	edit := widget.NewMultiLineEntry()

	edit.Wrapping = fyne.TextWrapWord
	edit.Resize(fyne.NewSize(600, 600))
	edit.SetText("")

	l_snips := widget.NewList(
		func() int {
			return 0
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {},
	)

	l_groups := widget.NewList(
		func() int {
			return len(app.Snips.Groups)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(idg widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[idg].Category)
		},
	)

	l_groups.OnSelected = app.refreshGroup
	l_snips.OnSelected = app.refreshSnip
	filter.OnChanged = app.refreshFilter
	edit.OnChanged = app.refreshEdit

	app.FilterWidget = filter
	app.Btn = btn
	app.EditWidget = edit
	app.ListSnip = l_snips
	app.ListGroup = l_groups

	return filter, btn, edit, l_snips, l_groups
}

func (app *config) makeUI_move(win fyne.Window) *widget.List {
	l_groups := widget.NewList(
		func() int {
			return len(app.Snips.Groups)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(idg widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[idg].Category)
		},
	)

	l_groups.OnSelected = app.refreshGroupInsSnip

	return l_groups
}

func (app *config) makeUIGrMove(win fyne.Window) *widget.List {
	lGroups := widget.NewList(
		func() int {
			return len(app.Snips.Groups)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(idg widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[idg].Category)
		},
	)

	lGroups.OnSelected = app.refreshGroupMove

	return lGroups
}

func (app *config) makeUI_groupIns(win fyne.Window) (*widget.Label, *widget.Entry, *widget.Button) {
	label := widget.NewLabel("Group name: ")
	edit := widget.NewEntry()
	edit.Resize(fyne.NewSize(50, 200))
	btn := widget.NewButton("Ok", func() {
		gr := Group{Category: edit.Text, Language: "PLSQL"}
		app.Snips.Groups = append(app.Snips.Groups, gr)
		app.IDGroup = app.IDGroup + 1
		app.NewGroupEdit.SetText("")
		app.resortGroups()
		app.ListGroup.Refresh()
		app.winGrIns.Hide()
	})

	app.NewGroupEdit = edit

	return label, edit, btn
}

func (app *config) makeUINewName(win fyne.Window) (*widget.Entry, *widget.Button) {

	edit := widget.NewEntry()
	edit.Resize(fyne.NewSize(50, 200))
	btn := widget.NewButton("Ok", func() {
		// TODO: if not nil
		switch app.WhoActive {
		case "Snip":
			{
				app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Name = edit.Text
				app.resortSnips()
				app.ListSnip.Refresh()
				app.ListSnip.Select(app.IDSnip)
			}
		case "Group":
			{
				app.Snips.Groups[app.IDGroup].Category = edit.Text
				app.resortGroups()
				app.ListGroup.Refresh()
				app.ListGroup.Select(app.IDGroup)
			}
		}
		app.winNewName.Hide()
	})

	app.NewNameEdit = edit

	return edit, btn

}

func (app *config) makeUI_snipIns(win fyne.Window) (*widget.List, *widget.Label, *widget.Entry, *widget.Label, *widget.Entry, *widget.Button) {
	l_groups := widget.NewList(
		func() int {
			return len(app.Snips.Groups)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(idg widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[idg].Category)
		},
	)
	l_groups.OnSelected = app.refreshGroupInsSnip

	labelGr := widget.NewLabel("New group name: ")
	editGr := widget.NewEntry()
	editGr.Resize(fyne.NewSize(50, 200))
	labelSn := widget.NewLabel("New snippet name: ")
	editSn := widget.NewEntry()
	editSn.Resize(fyne.NewSize(50, 200))
	btn := widget.NewButton("Ok", func() {
		if editSn.Text != "" {
			app.Snips.Groups[app.IDGroup].Snips = append(app.Snips.Groups[app.IDGroup].Snips, Snippet{Name: editSn.Text})
			app.resortSnips()
			if editGr.Text != "" {
				app.Snips.Groups = append(app.Snips.Groups, Group{Category: editGr.Text, Language: "PLSQL"})
				app.setGroupByName(editGr.Text)
				app.resortGroups()
			} else {
				app.IDGroup = app.IDGroupNewSnip
			}
			app.ListGroup.Refresh()
		}

		app.NewGroupEdit.SetText("")
		app.NewSnipEdit.SetText("")
		app.winSnIns.Hide()

	})

	app.NewGroupEdit = editGr
	app.NewSnipEdit = editSn

	return l_groups, labelGr, editGr, labelSn, editSn, btn
}

func (app *config) searchFilter(token string) {
	if token == "" {
		app.Snips = app.SnipsDefault
		app.IDGroup = 0
		app.IDSnip = 0
		app.ListSnip.UnselectAll()
		app.refreshGroup(0)
		app.ListGroup.Select(0)
		return
	}

	app.Snips = app.SnipsDefault
	var newSnippets Snippets
	newSnippets.XMLName = app.Snips.XMLName
	for gi, gel := range app.Snips.Groups {
		var newGidx int
		var newG Group
		newGidx = -1
		for si, sel := range app.Snips.Groups[gi].Snips {
			if strings.Contains(strings.ToLower((app.Snips.Groups[gi].Snips[si].Code)), strings.ToLower(token)) {

				if newGidx == -1 {
					newGidx = gi
					newG.XMLName = gel.XMLName
					newG.Category = gel.Category
					newG.Language = gel.Language
				}
				newG.Snips = append(newG.Snips, sel)
			}

		}
		if newGidx != -1 {
			newSnippets.Groups = append(newSnippets.Groups, newG)
		}
	}

	if len(newSnippets.Groups) != 0 {
		app.Snips = newSnippets
	} else {
		app.Snips = app.SnipsDefault
		app.FilterWidget.SetText("")
	}
	app.IDGroup = 0
	app.IDSnip = 0
	app.ListGroup.UnselectAll()
	app.ListSnip.UnselectAll()

	app.refreshGroup(0)
	app.ListGroup.Select(0)
}
