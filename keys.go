package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (app *config) setKeys(win fyne.Window) {
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		// fmt.Println(k.Name)
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
						app.ListSnip.UnselectAll()
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

func (app *config) setKeysMove(win fyne.Window) {
	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println(k.Name)
		switch k.Name {
		case "KP_Enter":
			{
				// TODO: Rewrite logic
				snip := app.Snips.Groups[app.IDGroup].Snips[app.IDSnip]
				down_range := app.IDSnip - 1
				if down_range < 0 {
					down_range = 1
				}
				up_range := app.IDSnip + 1
				if up_range > app.ListSnip.Length() {
					up_range = app.IDSnip - 1
				}

				left_arr, right_arr := app.Snips.Groups[app.IDGroup].Snips[down_range:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[up_range:]
				app.Snips.Groups[app.IDGroup].Snips = append(left_arr, right_arr...)
				app.Snips.Groups[app.IDGroupMove].Snips = append(app.Snips.Groups[app.IDGroupMove].Snips, snip)
				win.Hide()

			}
		case "Return":
			{
				// TODO: Rewrite logic
				snip := app.Snips.Groups[app.IDGroup].Snips[app.IDSnip]
				down_range := app.IDSnip - 1
				if down_range < 0 {
					down_range = 1
				}
				up_range := app.IDSnip + 1
				if up_range > app.ListSnip.Length() {
					up_range = app.IDSnip - 1
				}

				left_arr, right_arr := app.Snips.Groups[app.IDGroup].Snips[down_range:app.IDSnip], app.Snips.Groups[app.IDGroup].Snips[up_range:]
				app.Snips.Groups[app.IDGroup].Snips = append(left_arr, right_arr...)
				app.Snips.Groups[app.IDGroupMove].Snips = append(app.Snips.Groups[app.IDGroupMove].Snips, snip)
				win.Hide()
			}

		}
	})
}
