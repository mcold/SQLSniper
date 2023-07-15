package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
	WhoActive    string
}

var cfg config

func main() {
	a := app.New()

	win := a.NewWindow("Sniper")

	edit, l_snips, l_groups := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(l_groups, container.NewHSplit(l_snips, edit)))

	cfg.WhoActive = "Group"

	cfg.setArrows(win)

	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) setArrows(win fyne.Window) {
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case "Up":
			{
				if app.WhoActive == "Group" {
					newID := app.IDGroup - 1
					if newID >= 0 {
						app.IDGroup = newID
						app.ListGroup.Select(app.IDGroup)
						app.IDSnip = 0
						app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
						app.EditWidget.Refresh()
					}
				} else if app.WhoActive == "Snip" {
					newID := app.IDSnip - 1
					if newID >= 0 {
						app.IDSnip = newID
						app.ListSnip.Select(app.IDSnip)
						app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
						app.EditWidget.Refresh()
					}
				}
			}
		case "Down":
			{
				if app.WhoActive == "Group" {
					newID := app.IDGroup + 1
					if app.ListGroup.Length() > newID {
						app.IDGroup = newID
						app.IDSnip = 0
						app.ListGroup.Select(app.IDGroup)
						app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
						app.EditWidget.Refresh()
					}

				} else if app.WhoActive == "Snip" {
					newID := app.IDSnip + 1
					if app.ListSnip.Length() > newID {
						app.IDSnip = newID
						app.ListSnip.Select(app.IDSnip)
						app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
						app.EditWidget.Refresh()
					}
				}

			}
		case "Right":
			{

				switch app.WhoActive {
				case "Group":
					{
						app.WhoActive = "Snip"
						app.IDSnip = 0
						app.ListSnip.Select(app.IDSnip)
						app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
						app.EditWidget.Refresh()
					}
				case "Snip":
					{
						app.WhoActive = "Edit"
						app.EditWidget.SelectedText()
						app.EditWidget.Enable()
					}
				}

			}
		case "Left":
			{
				switch app.WhoActive {
				case "Snip":
					{
						app.WhoActive = "Group"
					}
				case "Edit":
					{
						app.WhoActive = "Snip"
					}
				}
			}

		}

	})
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
	// app.ListSnip.Select(0)
	app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
	// app.refreshSnip(app.IDSnip)
}

func (app *config) refreshSnip(id int) {
	app.WhoActive = "Snip"
	app.IDSnip = id
	app.EditWidget.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
	app.EditWidget.Refresh()
}

func (app *config) makeUI() (*widget.Entry, *widget.List, *widget.List) {

	app.Snips = get_snippets("UserSnippets.xml")

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

func (app *config) createMenuItems(win fyne.Window) {

	// create three menu items
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win))
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	// create a file menu, and add the three items to it
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// create a main menu, and add the file menu to it
	menu := fyne.NewMainMenu(fileMenu)

	// set the main menu for the application
	win.SetMainMenu(menu)
}

var filter = storage.NewExtensionFileFilter([]string{".xml", ".XML"})

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			defer write.Close()
		}
	}
}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))

			app.CurrentFile = read.URI()
			app.Snips = get_snippets(app.CurrentFile.Name())
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			app.SaveMenuItem.Disabled = false

			fmt.Println(app.CurrentFile.Name())

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".xml") {
				dialog.ShowInformation("Error", "Please name your file with a .xml extension!", win)
				return
			}

			// save file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)
		saveDialog.SetFileName("UserSnippets.xml")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
