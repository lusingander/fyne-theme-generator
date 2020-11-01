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

	connected bool
	toolbar   *widget.Toolbar
}

func (u *ui) newWidgetsPanel() *widgetsPanel {
	p := &widgetsPanel{
		parent:    u.window,
		connected: true,
	}
	p.build(u.connectWidgetPanel, u.disconnectWidgetPanel)
	return p
}

func (p *widgetsPanel) connect(onClickDisconnect func()) {
	p.connected = true
	p.resetToolbar(nil, onClickDisconnect)
}

func (p *widgetsPanel) disconnect(onClickConnect func()) {
	p.connected = false
	p.resetToolbar(onClickConnect, nil)
}

func (p *widgetsPanel) build(onClickConnect, onClickDisconnect func()) {
	p.resetToolbar(onClickConnect, onClickDisconnect)
	p.panel = fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		p.toolbar,
		p.labels(),
		p.buttons(),
		p.inputs(),
		p.progresses(),
	)
}

func (p *widgetsPanel) resetToolbar(onClickConnect, onClickDisconnect func()) {
	var action widget.ToolbarItem
	if p.connected {
		action = widget.NewToolbarAction(theme.ContentAddIcon(), onClickDisconnect)
	} else {
		action = widget.NewToolbarAction(theme.ContentClearIcon(), onClickConnect)
	}
	if p.toolbar == nil {
		p.toolbar = widget.NewToolbar(
			widget.NewToolbarSpacer(),
			newToolbarLabel("Fyne Theme Generator"),
			widget.NewToolbarSpacer(),
			action,
		)
	} else {
		// HACK: toolbarRenderer#resetObjects is not called when len(Items) is not changed...
		p.toolbar.Items = p.toolbar.Items[:len(p.toolbar.Items)-1]
		p.toolbar.Refresh()
		p.toolbar.Append(action)
	}
}

type toolbarLabel struct {
	fyne.CanvasObject
}

func newToolbarLabel(label string) *toolbarLabel {
	align := fyne.TextAlignCenter
	style := fyne.TextStyle{Bold: true}
	return &toolbarLabel{widget.NewLabelWithStyle(label, align, style)}
}

func (l *toolbarLabel) ToolbarObject() fyne.CanvasObject {
	return l.CanvasObject
}

func (*widgetsPanel) labels() fyne.CanvasObject {
	parsed, _ := url.Parse(repositoryURL)
	link := widget.NewHyperlink("repository", parsed)
	description := widget.NewLabel("WYSIWYG theme editor for Fyne")
	description.TextStyle.Italic = true
	bold := widget.NewLabel("Bold text")
	bold.TextStyle.Bold = true
	boldItalic := widget.NewLabel("Bold Italic text")
	boldItalic.TextStyle.Bold = true
	boldItalic.TextStyle.Italic = true
	monospace := widget.NewLabel("Monospace text")
	monospace.TextStyle.Monospace = true
	return widget.NewGroup(
		"Labels",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				description,
				link,
			),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				bold,
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
