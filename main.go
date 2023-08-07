package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget        *widget.Entry
	EditToolTip       *widget.Entry
	FilterWidget      *widget.Entry
	FilterMoveGr      *widget.Entry
	Btn               *widget.Button
	ListSnip          *widget.List
	ListGroup         *widget.List
	ListGroupMoveToGr *widget.List
	CurrentFile       fyne.URI
	PathFile          string
	SaveMenuItem      *fyne.MenuItem
	Snips             Snippets
	SnipsDefault      Snippets
	GroupsDefault     []Group
	IDGroup           int
	IDGroupMove       int
	IDGroupSnipMove   int
	IDGroupNewSnip    int
	IDSnip            int
	WhoActive         string
	winMove           fyne.Window
	winGrMove         fyne.Window
	winGrIns          fyne.Window
	winSnIns          fyne.Window
	winNewName        fyne.Window
	NewGroupEdit      *widget.Entry
	NewSnipEdit       *widget.Entry
	NewNameEdit       *widget.Entry
}

var cfg config

func main() {
	a := app.New()

	win := a.NewWindow("Sniper")

	winMove := a.NewWindow("Move Snippet")
	l_groups_move := cfg.makeUI_move(winMove)
	winMove.Resize(fyne.Size{Width: 320, Height: 400})
	winMove.SetContent(l_groups_move)

	cfg.winMove = winMove

	winGrMove := a.NewWindow("Group move")
	l_group_move, filterGrMove, btnFilterGrMove := cfg.makeUIGrMove(winGrMove)
	winGrMove.Resize(fyne.Size{Width: 320, Height: 400})
	winGrMove.SetContent(container.NewVSplit(container.NewVBox(filterGrMove, btnFilterGrMove), l_group_move))

	cfg.winGrMove = winGrMove

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

	filter, btn, edit, l_snips, l_groups, editToolTip := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(container.NewVSplit(l_groups, container.NewVBox(filter, btn)), container.NewHSplit(l_snips, container.NewVSplit(editToolTip, container.NewHScroll(edit)))))

	winNewGroupName := a.NewWindow("New name")
	NewNameEdit, NewNameBtn := cfg.makeUINewName(winNewGroupName)
	winNewGroupName.Resize(fyne.Size{Width: 300, Height: 100})
	NewNameEdit.Resize(fyne.Size{Width: 200, Height: 50})
	winNewGroupName.SetContent(container.NewVBox(NewNameEdit, NewNameBtn))

	cfg.winNewName = winNewGroupName

	cfg.WhoActive = "Group"

	cfg.setKeys(win)
	cfg.setKeysMove(winMove)
	cfg.setKeysGrMove(winGrMove)

	cfg.SnipsDefault = cfg.Snips
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}
