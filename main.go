package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

type config struct {
	Snips         Snippets
	ListGrData    binding.ExternalStringList
	ListSnData    binding.ExternalStringList
	SnipsArr      []string
	EditSnip      string
	EditSnipData  binding.ExternalString
	EditDescr     string
	EditDescrData binding.ExternalString
	IDGroup       int
	IDSnip        int
}

var cfg config

func main() {
	a := app.New()

	win := a.NewWindow("Sniper")

	listGr, listSn, EditSn, EditDescr := cfg.makeUI()

	listGr.Resize(fyne.Size{Width: 150, Height: 500})
	listGr.Move(fyne.Position{X: 0, Y: 0})

	listSn.Resize(fyne.Size{Width: 150, Height: 500})
	listSn.Move(fyne.Position{X: 150, Y: 0})

	EditSn.Resize(fyne.Size{Width: 500, Height: 500})
	EditSn.Move(fyne.Position{X: 300, Y: 75})

	EditDescr.Resize(fyne.Size{Width: 500, Height: 75})
	EditDescr.Move(fyne.Position{X: 300, Y: 0})

	c1 := container.NewWithoutLayout(listGr, listSn, EditSn, EditDescr)

	win.SetContent(c1)

	win.Resize(fyne.Size{Width: 800, Height: 500})

	win.ShowAndRun()
}
