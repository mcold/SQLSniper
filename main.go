package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget     *widget.Entry
	FilterWidget   *widget.Entry
	Btn            *widget.Button
	ListSnip       *widget.List
	ListGroup      *widget.List
	CurrentFile    fyne.URI
	PathFile       string
	SaveMenuItem   *fyne.MenuItem
	Snips          Snippets
	SnipsDefault   Snippets
	IDGroup        int
	IDGroupMove    int
	IDGroupNewSnip int
	IDSnip         int
	WhoActive      string
	winMove        fyne.Window
	winGrIns       fyne.Window
	winSnIns       fyne.Window
	winNewName     fyne.Window
	NewGroupEdit   *widget.Entry
	NewSnipEdit    *widget.Entry
	NewNameEdit    *widget.Entry
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
	winMove.SetContent(l_groups_move)

	cfg.winMove = winMove

	winGroupIns := a.NewWindow("New Group")
	GrLabel, GrEdit, GrBtn := cfg.makeUI_groupIns(winGroupIns)
	winGroupIns.Resize(fyne.Size{Width: 300, Height: 100})
	GrEdit.Resize(fyne.Size{Width: 200, Height: 50})
	winGroupIns.SetContent(container.NewVBox(container.NewHBox(GrLabel, GrEdit), GrBtn))

	cfg.winGrIns = winGroupIns

	winSnipIns := a.NewWindow("New Snippet")
	GrSnList, GrSnLabel, GrSnEdit, SnLabel, SnEdit, SnBtn := cfg.makeUI_snipIns(winSnipIns)
	winSnipIns.Resize(fyne.Size{Width: 300, Height: 200})
	GrSnEdit.Resize(fyne.Size{Width: 200, Height: 50})
	SnEdit.Resize(fyne.Size{Width: 200, Height: 50})
	winGroupIns.SetContent(container.NewVBox(GrSnList, container.NewHBox(GrSnLabel, GrSnEdit), container.NewHBox(SnLabel, SnEdit), SnBtn))

	cfg.winSnIns = winSnipIns

	// cfg.winMove.CenterOnScreen()

	filter, btn, edit, l_snips, l_groups := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(container.NewVSplit(l_groups, container.NewVBox(filter, btn)), container.NewHSplit(l_snips, edit)))

	winNewGroupName := a.NewWindow("New name")
	NewNameEdit, NewNameBtn := cfg.makeUINewName(winNewGroupName)
	winNewGroupName.Resize(fyne.Size{Width: 300, Height: 100})
	NewNameEdit.Resize(fyne.Size{Width: 200, Height: 50})
	winNewGroupName.SetContent(container.NewVBox(NewNameEdit, NewNameBtn))

	cfg.winNewName = winNewGroupName

	cfg.WhoActive = "Group"

	cfg.setKeys(win)
	cfg.setKeysMove(winMove)

	cfg.SnipsDefault = cfg.Snips
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}
