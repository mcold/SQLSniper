package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (app *config) refreshFilter(token string) {
	app.Snips = app.SnipsDefault
	app.WhoActive = "Filter"
}

func (app *config) refreshSnip(id int) {
	app.WhoActive = "Snip"
	app.IDSnip = id
	app.EditToolTip.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Description)
	app.EditToolTip.Refresh()
	app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
	app.EditWidget.Refresh()
}

func (app *config) refreshGroup(id int) {
	app.WhoActive = "Group"
	app.IDGroup = id
	app.ListSnip.Length = func() int { return len(app.Snips.Groups[id].Snips) }
	app.ListSnip.UpdateItem = func(ids widget.ListItemID, obj fyne.CanvasObject) {
		obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[id].Snips[ids].Name)
	}
	app.ListSnip.Refresh()

	// activate first element by default
	app.IDSnip = 0
	app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
}

func (app *config) refreshGroupMove(id int) {
	app.WhoActive = "GroupMove"
	app.IDGroupMove = id
}

func (app *config) refreshGroupInsSnip(id int) {
	app.WhoActive = "GroupInsSnip"
	app.IDGroupNewSnip = id
}

func (app *config) refreshGroupMoveSnip(id int) {
	app.WhoActive = "GroupMoveSnip"
	app.IDGroupSnipMove = id
}

func (app *config) refreshEditToolTip(str string) {
	app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Description = str
}

func (app *config) refreshEdit(str string) {
	app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code = str
}

func (app *config) getGroupByName(str string) int {
	var out_var int
	for idx, el := range app.Snips.Groups {
		if el.Category == str {
			out_var = idx
			break
		}
	}
	return out_var
}

func (app *config) setGroupByName(str string) {
	for idx, el := range app.Snips.Groups {
		if el.Category == str {
			app.IDGroup = idx
			break
		}
	}
}
