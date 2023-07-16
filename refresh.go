package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (app *config) refreshFilter(token string) {
	app.WhoActive = "Filter"
}

func (app *config) refreshSnip(id int) {
	app.WhoActive = "Snip"
	app.IDSnip = id
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

func (app *config) refreshEdit(str string) {
	app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code = str
}
