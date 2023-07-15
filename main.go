package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget   *widget.Entry
	ListSnip     *widget.List
	ListGroup    *widget.List
	CurrentFile  fyne.URI
	SaveMenuItem *fyne.MenuItem
	Snips        Snippets
	IDGroup      int
	IDSnip       int
}

var cfg config

func main() {
	a := app.New()

	win := a.NewWindow("Sniper")

	edit, l_snips, l_groups := cfg.makeUI()

	// set the content of the window
	win.SetContent(container.NewHSplit(l_groups, container.NewHSplit(l_snips, edit)))

	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) refreshGroup(id int) {
	app.IDGroup = id
	app.ListSnip.Length = func() int { return len(app.Snips.Groups[id].Snips) }
	app.ListSnip.UpdateItem = func(ids widget.ListItemID, obj fyne.CanvasObject) {
		obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[id].Snips[ids].Name)
	}
	app.ListSnip.Refresh()

	// activate first element by default
	app.IDSnip = 0
	app.ListSnip.Select(0)
	app.refreshSnip(app.IDSnip)
}

func (app *config) refreshSnip(id int) {
	app.IDSnip = id
	app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
	app.EditWidget.Refresh()

}

func (app *config) makeUI() (*widget.Entry, *widget.List, *widget.List) {

	app.Snips = get_snippets()

	edit := widget.NewMultiLineEntry()
	edit.Wrapping = fyne.TextWrapWord
	edit.Resize(fyne.NewSize(600, 600))
	edit.SetText("")

	l_snips := widget.NewList(
		func() int {
			// return cnt_snips
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

	app.EditWidget = edit
	app.ListSnip = l_snips
	app.ListGroup = l_groups

	return edit, l_snips, l_groups
}
