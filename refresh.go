package main

func (app *config) refreshGroup(id int) {
	app.IDGroup = id
	snips := []string{}

	for si := range app.Snips.Groups[id].Snips {
		snips = append(snips, app.Snips.Groups[id].Snips[si].Name)
	}

	app.SnipsArr = snips
	app.ListSnData.Reload()
}

func (app *config) refreshSnip(id int) {

	app.IDSnip = id

	app.EditSnipData.Set(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Code)
	app.EditSnipData.Reload()

	app.EditDescrData.Set(app.Snips.Groups[app.IDGroup].Snips[app.IDSnip].Description)
	app.EditDescrData.Reload()
}
