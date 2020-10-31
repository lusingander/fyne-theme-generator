package ui

import (
	"net/url"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/data/validation"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	repositoryURL = "https://github.com/lusingander/fyne-theme-generator"
)

type widgetsPanel struct {
	panel  fyne.CanvasObject
	parent fyne.Window
}

func (u *ui) newWidgetsPanel() *widgetsPanel {
	p := &widgetsPanel{
		parent: u.window,
	}
	p.build()
	return p
}

func (p *widgetsPanel) build() {
	p.panel = fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		p.labels(),
		p.buttons(),
		p.inputs(),
		p.progresses(),
	)
}

func (*widgetsPanel) labels() fyne.CanvasObject {
	title := widget.NewLabel("Fyne Theme Generator")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	parsed, _ := url.Parse(repositoryURL)
	link := widget.NewHyperlink("repository", parsed)
	description := widget.NewLabel("WYSIWYG theme editor for Fyne")
	description.TextStyle.Italic = true
	boldItalic := widget.NewLabel("Bold Italic text")
	boldItalic.TextStyle.Bold = true
	boldItalic.TextStyle.Italic = true
	monospace := widget.NewLabel("Monospace text")
	monospace.TextStyle.Monospace = true
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
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				boldItalic,
				monospace,
			),
		),
	)
}

func (p *widgetsPanel) buttons() fyne.CanvasObject {
	info := widget.NewButton("Info", func() {
		dialog.NewInformation("Info", "information dialog...", p.parent)
	})
	confirm := widget.NewButton("Confirm", func() {
		dialog.NewConfirm("Confirm", "confirm dialog...", func(bool) {}, p.parent)
	})
	file := widget.NewButtonWithIcon("File Dialog", theme.FolderOpenIcon(), func() {
		dialog.ShowFileOpen(func(fyne.URIReadCloser, error) {}, p.parent)
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

func (*widgetsPanel) inputs() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")
	disabledEntry := widget.NewEntry()
	disabledEntry.Disable()
	disabledEntry.SetText("Entry (disabled)")
	validateEntry := widget.NewEntry()
	validateEntry.SetPlaceHolder("Must contain a number")
	validateEntry.SetText("abc")
	validateEntry.Validator = validation.NewRegexp(`\d`, "Must contain a number")

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
			validateEntry,
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

func (*widgetsPanel) progresses() fyne.CanvasObject {
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
