package ui

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	p.build(u.disconnectWidgetPanel)
	return p
}

func (p *widgetsPanel) connect(onClickDisconnect func()) {
	p.connected = true
	p.resetToolbar(theme.ContentAddIcon(), onClickDisconnect)
}

func (p *widgetsPanel) disconnect(onClickConnect func()) {
	p.connected = false
	p.resetToolbar(theme.ContentClearIcon(), onClickConnect)
}

func (p *widgetsPanel) build(onClickDisconnect func()) {
	p.connect(onClickDisconnect)
	p.panel = fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		p.toolbar,
		p.labels(),
		p.buttons(),
		p.inputs(),
		p.progresses(),
		p.others(),
	)
}

func (p *widgetsPanel) resetToolbar(icon fyne.Resource, onClick func()) {
	action := widget.NewToolbarAction(icon, onClick)
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
	return widget.NewCard(
		"Labels",
		"",
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
		dialog.ShowInformation("Info", "information dialog...", p.parent)
	})
	confirm := widget.NewButton("Confirm", func() {
		dialog.ShowConfirm("Confirm", "confirm dialog...", func(bool) {}, p.parent)
	})
	file := widget.NewButtonWithIcon("File Dialog", theme.FolderOpenIcon(), func() {
		dialog.ShowFileOpen(func(fyne.URIReadCloser, error) {}, p.parent)
	})
	disabled := widget.NewButtonWithIcon("Disabled", theme.CancelIcon(), func() {})
	disabled.Disable()
	return widget.NewCard(
		"Buttons & Dialogs",
		"",
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
	// validateEntry.SetText("abc") // not working expectedly in v2...
	validateEntry.Validator = validation.NewRegexp(`\d`, "Must contain a number")

	selects := widget.NewSelect([]string{"Foo", "Bar", "Baz"}, func(string) {})
	check := widget.NewCheck("Check", func(bool) {})
	disabledCheck := widget.NewCheck("Check (disabled)", func(bool) {})
	disabledCheck.Disable()
	radio := widget.NewRadioGroup([]string{"Radio"}, func(string) {})
	disabledRadio := widget.NewRadioGroup([]string{"Radio (disabled)"}, func(string) {})
	disabledRadio.Disable()
	slider := widget.NewSlider(0, 100)
	return widget.NewCard(
		"Inputs",
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				entry,
				disabledEntry,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(2),
				validateEntry,
				selects,
			),
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
	// changes are not reflected immediately...
	// progressInf := widget.NewProgressBarInfinite()
	return widget.NewCard(
		"Progresses",
		"",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			progress,
			// progressInf,
		),
	)
}

func (*widgetsPanel) others() fyne.CanvasObject {
	creditsButton := widget.NewButton("CREDITS", func() {
		CreditsWindow(fyne.CurrentApp(), fyne.NewSize(800, 600)).Show()
	})
	return widget.NewCard(
		"Others",
		"",
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			creditsButton,
		),
	)
}
