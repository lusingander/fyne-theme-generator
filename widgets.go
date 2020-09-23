package main

import (
	"net/url"

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
