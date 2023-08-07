package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
)

func (app *config) setKeys(win fyne.Window) {
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println(k.Name)
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
						app.EditWidget.SelectedText()
					}
				case "Snip":
					{
						app.WhoActive = "Edit"
						app.EditWidget.Enable()
						app.EditWidget.Refresh()
						// fmt.Println(app.EditWidget.CursorRow)

					}
				}

			}
		case "Left":
			{
				switch app.WhoActive {
				case "Snip":
					{
						app.WhoActive = "Group"
						app.ListSnip.UnselectAll()
					}
				case "Edit":
					{
						app.WhoActive = "Snip"
					}
				}
			}
		case "Delete":
			{
				switch app.WhoActive {
				case "Snip":
					{
						if app.IDSnip+1 > len(app.Snips.Groups[app.IDGroup].Snips) {
							app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[:app.IDSnip]
						} else {
							app.Snips.Groups[app.IDGroup].Snips = append(app.Snips.Groups[app.IDGroup].Snips[:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[app.IDSnip+1:]...)
						}

						app.ListSnip.Refresh()
						app.EditWidget.Refresh()
					}
				case "Edit":
					{
						if app.IDSnip+1 > len(app.Snips.Groups[app.IDGroup].Snips) {
							app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[:app.IDSnip]
						} else {
							app.Snips.Groups[app.IDGroup].Snips = append(app.Snips.Groups[app.IDGroup].Snips[:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[app.IDSnip+1:]...)
						}
						app.ListSnip.Refresh()
						app.EditWidget.Refresh()
					}
				case "Group":
					{
						if app.IDGroup+1 > app.ListGroup.Length() {
							app.Snips.Groups = app.Snips.Groups[:app.IDGroup]

							app.IDGroup = app.IDGroup - 1
							app.ListGroup.Select(app.IDGroup)
						} else {
							app.Snips.Groups = append(app.Snips.Groups[:app.IDGroup], app.Snips.Groups[app.IDGroup+1:]...)

							app.IDGroup = app.IDGroup - 1
							app.ListGroup.Select(app.IDGroup)
						}

						app.ListGroup.Refresh()
						app.ListSnip.Refresh()
						app.EditWidget.Refresh()
					}
				}

			}
		case "Insert":
			{
				switch app.WhoActive {
				case "Snip":
					{
						app.winSnIns.CenterOnScreen()
						app.winSnIns.Show()
					}
				case "Group":
					{
						app.winGrIns.CenterOnScreen()
						app.winGrIns.Show()

					}
				}
			}
		case "F2":
			{
				switch app.WhoActive {
				case "Snip":
					{
						//fmt.Println(app.Snips.Groups[app.IDGroup].Category)
						app.NewNameEdit.SetText(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Name)
						app.winNewName.CenterOnScreen()
						app.winNewName.Show()
					}
				case "Group":
					{
						app.NewNameEdit.SetText(app.Snips.Groups[app.IDGroup].Category)
						app.winNewName.CenterOnScreen()
						app.winNewName.Show()
					}
				}
			}
		}
	},
	)
}

func (app *config) setKeysMove(win fyne.Window) {
	// TODO: not show current group
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println(k.Name)
		switch k.Name {
		case "KP_Enter":
			{
				snip := app.Snips.Groups[app.IDGroup].Snips[app.IDSnip]
				down_range := app.IDSnip - 1
				up_range := app.IDSnip + 1

				fmt.Println("down_range: " + strconv.Itoa(down_range))
				fmt.Println("up_range: " + strconv.Itoa(up_range))

				if down_range < 0 {
					if up_range > len(app.Snips.Groups[app.IDGroup].Snips) {
						// fmt.Println("THE END!!!")
						if app.IDGroup+1 > app.ListGroup.Length() {
							app.Snips.Groups = app.Snips.Groups[app.IDGroup:]
							app.IDGroupSnipMove = app.IDGroupSnipMove - 1
						} else {
							left_group_arr := app.Snips.Groups[:app.IDGroup]
							right_group_arr := app.Snips.Groups[app.IDGroup+1:]
							app.Snips.Groups = append(left_group_arr, right_group_arr...)
						}

						app.IDGroup = 0
					} else {
						app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[up_range:]
					}
				} else {
					if up_range > len(app.Snips.Groups[app.IDGroup].Snips) {
						app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[down_range:app.IDSnip]
					} else {
						left_arr, right_arr := app.Snips.Groups[app.IDGroup].Snips[:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[up_range:]
						app.Snips.Groups[app.IDGroup].Snips = append(left_arr, right_arr...)
					}
				}
				fmt.Println("IDGroupSnipMove: " + strconv.Itoa(app.IDGroupSnipMove))
				app.Snips.Groups[app.IDGroupSnipMove].Snips = append(app.Snips.Groups[app.IDGroupSnipMove].Snips, snip)
				app.refreshGroup(app.IDGroupSnipMove)
				app.IDGroup = app.IDGroupSnipMove
				app.ListSnip.UnselectAll()
				app.ListGroup.Select(app.IDGroup)
				app.IDSnip = len(app.Snips.Groups[app.IDGroup].Snips) - 1
				// app.IDSnip = 0
				app.ListSnip.Select(app.IDSnip)
				// app.ListGroup.Refresh()
				// app.ListSnip.Refresh()
				win.Hide()

			}
		case "Return":
			{
				// TODO: Rewrite logic
				snip := app.Snips.Groups[app.IDGroup].Snips[app.IDSnip]
				down_range := app.IDSnip - 1
				up_range := app.IDSnip + 1

				fmt.Println("down_range: " + strconv.Itoa(down_range))
				fmt.Println("up_range: " + strconv.Itoa(up_range))

				if down_range < 0 {
					if up_range > app.ListSnip.Length() {
						if app.IDGroup+1 > app.ListGroup.Length() {
							app.Snips.Groups = app.Snips.Groups[app.IDGroup:]
							app.IDGroupMove = app.IDGroupMove - 1
							app.ListGroup.Refresh()
						} else {
							left_group_arr := app.Snips.Groups[:app.IDGroup]
							right_group_arr := app.Snips.Groups[app.IDGroup+1:]
							app.Snips.Groups = append(left_group_arr, right_group_arr...)
						}

						app.IDGroup = 0
					} else {
						app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[up_range:]
					}
				} else {
					if up_range > app.ListSnip.Length() {
						app.Snips.Groups[app.IDGroup].Snips = app.Snips.Groups[app.IDGroup].Snips[down_range:app.IDSnip]
					} else {
						left_arr, right_arr := app.Snips.Groups[app.IDGroup].Snips[down_range:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[up_range:]
						app.Snips.Groups[app.IDGroup].Snips = append(left_arr, right_arr...)
					}
				}

				app.Snips.Groups[app.IDGroupMove].Snips = append(app.Snips.Groups[app.IDGroupMove].Snips, snip)
				app.refreshGroup(app.IDGroupMove)
				app.ListGroup.Refresh()
				win.Hide()
			}

		}
	})
}

func (app *config) setKeysGrMove(win fyne.Window) {
	// TODO: not show current group
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println(k.Name)
		switch k.Name {
		case "KP_Enter":
			{
				app.Snips.Groups = app.GroupsDefault

				for i := range app.Snips.Groups[app.IDGroup].Snips {
					app.Snips.Groups[app.IDGroupMove].Snips = append(app.Snips.Groups[app.IDGroupMove].Snips, app.Snips.Groups[app.IDGroup].Snips[i])
				}

				groups := make([]Group, 0)
				cat_move_name := app.Snips.Groups[app.IDGroupMove].Category

				for i := range app.Snips.Groups {
					if app.Snips.Groups[i].Category != app.Snips.Groups[app.IDGroup].Category {
						groups = append(groups, app.Snips.Groups[i])
					}
				}

				fmt.Println(len(groups))
				app.Snips.Groups = groups

				for i := range app.Snips.Groups {
					if app.Snips.Groups[i].Category == cat_move_name {
						app.IDGroup = i
						break
					}
				}
				app.ListGroup.Refresh()
				app.ListSnip.UnselectAll()
				app.ListSnip.Refresh()
				win.Hide()

			}
		case "Return":
			{
				for i := range app.Snips.Groups[app.IDGroup].Snips {
					app.Snips.Groups[app.IDGroupMove].Snips = append(app.Snips.Groups[app.IDGroupMove].Snips, app.Snips.Groups[app.IDGroup].Snips[i])
				}

				groups := make([]Group, 0)
				cat_move_name := app.Snips.Groups[app.IDGroupMove].Category

				for i := range app.Snips.Groups {
					if app.Snips.Groups[i].Category != app.Snips.Groups[app.IDGroup].Category {
						groups = append(groups, app.Snips.Groups[i])
					}
				}

				fmt.Println(len(groups))
				app.Snips.Groups = groups

				for i := range app.Snips.Groups {
					if app.Snips.Groups[i].Category == cat_move_name {
						app.IDGroup = i
						break
					}
				}
				app.ListGroup.Refresh()
				app.ListSnip.UnselectAll()
				app.ListSnip.Refresh()
				win.Hide()
			}

		}
	})
}
