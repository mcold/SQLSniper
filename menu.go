package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func (app *config) createMenuItems(win fyne.Window) {

	// create file menu items
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win))
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	// create edit menu items
	moveMenuItem := fyne.NewMenuItem("Move...", app.moveSnip(win))

	// create a file menu, and add the three items to it
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	editMenu := fyne.NewMenu("Edit", moveMenuItem)

	// create a main menu, and add the file menu to it
	menu := fyne.NewMainMenu(fileMenu, editMenu)

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

			// save application config

			// read config
			// if _, err := toml.DecodeFile("conf.toml", &appCfg); err != nil {
			// 	fmt.Println("Error!")
			// }

			// fmt.Println(app.CurrentFile.Name())

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
			// write.Write([]byte(app.EditWidget.Text)) // only one snip

			///////////////// write xml
			e := xml.NewEncoder(write)

			// Write marshal the sninps struct to the file
			err = e.Encode(app.Snips)
			if err != nil {
				panic(err)
			}
			////////////////////////

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

func (app *config) moveSnip(win fyne.Window) func() {
	return func() {
		app.winMove.CenterOnScreen()
		app.winMove.Show()
	}
}
