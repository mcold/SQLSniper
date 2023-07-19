package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget   *widget.Entry
	FilterWidget *widget.Entry
	Btn          *widget.Button
	ListSnip     *widget.List
	ListGroup    *widget.List
	CurrentFile  fyne.URI
	PathFile     string
	SaveMenuItem *fyne.MenuItem
	Snips        Snippets
	SnipsDefault Snippets
	IDGroup      int
	IDGroupMove  int
	IDSnip       int
	WhoActive    string
	winMove      fyne.Window
}

// type appConfig struct {
// 	FileURI string
// 	Title   string
// }

var cfg config

// var appCfg appConfig

func main() {
	a := app.New()

	// appCfg, err := toml.LoadFile("config.toml")

	// if _, err := toml.DecodeFile("conf.toml", &appCfg); err != nil {
	// 	fmt.Println("Error!")
	// }

	// if err != nil {
	// 	fmt.Println("Error ", err.Error())
	// }

	// cfg.CurrentFile = appCfg.FileURI

	win := a.NewWindow("Sniper")

	winMove := a.NewWindow("Move Snippet")
	l_groups_move := cfg.makeUI_move(winMove)
	winMove.Resize(fyne.Size{Width: 320, Height: 400})
	// winMove.SetContent(container.NewVSplit(l_groups_move, widget.NewLabel("some")))
	winMove.SetContent(l_groups_move)

	cfg.winMove = winMove
	// cfg.winMove.CenterOnScreen()

	filter, btn, edit, l_snips, l_groups := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(container.NewVSplit(l_groups, container.NewVBox(filter, btn)), container.NewHSplit(l_snips, edit)))

	cfg.WhoActive = "Group"

	cfg.setKeys(win)
	cfg.setKeysMove(winMove)

	cfg.SnipsDefault = cfg.Snips
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}
