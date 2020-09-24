package main

import (
	"net/url"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	repositoryURL = "https://github.com/lusingander/fyne-theme-generator"
)

func widgets() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		labels(),
		buttons(),
		inputs(),
		progresses(),
	)
}

func labels() fyne.CanvasObject {
	title := widget.NewLabel("Fyne Theme Generator")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	parsed, _ := url.Parse(repositoryURL)
	link := widget.NewHyperlink("repository", parsed)
	link.Alignment = fyne.TextAlignCenter
	description := widget.NewLabel("WYSIWYG theme editor for Fyne")
	description.Alignment = fyne.TextAlignCenter
	description.TextStyle.Italic = true
	return widget.NewGroup(
		"Labels",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			title,
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				description,
				link,
			),
		),
	)
}

func buttons() fyne.CanvasObject {
	info := widget.NewButton("Info", func() {
		dialog.NewInformation("Info", "information dialog...", mainWindow)
	})
	confirm := widget.NewButton("Confirm", func() {
		dialog.NewConfirm("Confirm", "confirm dialog...", func(bool) {}, mainWindow)
	})
	file := widget.NewButtonWithIcon("File Dialog", theme.FolderOpenIcon(), func() {
		dialog.ShowFileOpen(func(fyne.URIReadCloser, error) {}, mainWindow)
	})
	disabled := widget.NewButtonWithIcon("Disabled", theme.CancelIcon(), func() {})
	disabled.Disable()
	return widget.NewGroup(
		"Buttons & Dialogs",
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			info,
			confirm,
			file,
			disabled,
		),
	)
}

func inputs() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")
	disabledEntry := widget.NewEntry()
	disabledEntry.Disable()
	disabledEntry.SetText("Entry (disabled)")
	selects := widget.NewSelect([]string{"Foo", "Bar", "Baz"}, func(string) {})
	check := widget.NewCheck("Check", func(bool) {})
	disabledCheck := widget.NewCheck("Check (disabled)", func(bool) {})
	disabledCheck.Disable()
	radio := widget.NewRadio([]string{"Radio"}, func(string) {})
	disabledRadio := widget.NewRadio([]string{"Radio (disabled)"}, func(string) {})
	disabledRadio.Disable()
	slider := widget.NewSlider(0, 100)
	return widget.NewGroup(
		"Inputs",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				entry,
				disabledEntry,
			),
			selects,
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				check,
				disabledCheck,
				radio,
				disabledRadio,
			),
			slider,
		),
	)
}

func progresses() fyne.CanvasObject {
	progress := widget.NewProgressBar()
	go func() {
		var v float64
		for {
			v += 0.01
			progress.SetValue(v)
			time.Sleep(time.Millisecond * 100)
			if v >= 1 {
				v = 0
			}
		}
	}()
	progressInf := widget.NewProgressBarInfinite()
	return widget.NewGroup(
		"Progresses",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			progress,
			progressInf,
		),
	)
}
