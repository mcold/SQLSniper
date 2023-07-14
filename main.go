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
}

var cfg config

func main() {
	a := app.New()

	win := a.NewWindow("Sniper")
	// win.Resize(fyne.NewSize(600, 600))

	edit, l_snips, l_groups := cfg.makeUI()

	// set the content of the window
	win.SetContent(container.NewHSplit(l_groups, container.NewHSplit(l_snips, edit)))

	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) refresh(id int) {
	app.ListSnip.Length = func() int { return len(app.Snips.Groups[id].Snips) }
	app.ListSnip.UpdateItem = func(ids widget.ListItemID, obj fyne.CanvasObject) {
		obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[id].Snips[ids].Name)
	}
	app.ListSnip.Refresh()
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
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {},
	)

	l_groups := widget.NewList(
		func() int {
			return len(app.Snips.Groups)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(idg widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(app.Snips.Groups[idg].Category)

			// l_snips.Length = func() int {
			// 	return len(snippets.Groups[idg].Snippets)
			// }

			//l_snips.UpdateItem = func(ids widget.ListItemID, obj fyne.CanvasObject) {
			// obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(snippets.Groups[idg].Snippets[ids].Name)

			// etr_code.SetText(snippets.Groups[idg].Snippets[ids].Code)

			// cnt = cnt + 1
			// // str := "idg: %s", idg
			// etr_code.SetText(fmt.Sprintf("idg: %s\nids: %s\ncnt: %s\n\ncode: %s", idg, ids, cnt, snippets.Groups[idg].Snippets[ids].Code))
			// fmt.Println("idg: %s", idg)
			// fmt.Println("l_groups:")
			// fmt.Println("idg: ", idg)
			// fmt.Println("ids: ", ids)
			// fmt.Println(snippets.Groups[idg].Snippets[ids].Code)
			// fmt.Println("cnt: ", cnt)
			// fmt.Println()

			//}
			// l_snips.UnselectAll()
			// l_snips.Select(0)

			//l_snips.Resize(fyne.NewSize(200, 600))
			// l_snips.Show()
			// etr_code.Show()

		},
	)

	l_groups.OnSelected = app.refresh

	app.EditWidget = edit
	app.ListSnip = l_snips
	app.ListGroup = l_groups

	return edit, l_snips, l_groups

	// app.PreviewWidget = preview

	// edit.OnChanged = preview.ParseMarkdown

	// return edit, preview
}
