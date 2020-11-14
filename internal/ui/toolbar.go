package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

type toolbarPanel struct {
	panel   fyne.CanvasObject
	parent  fyne.Window
	current *theme.Setting
}

func (u *ui) newToolbarPanel() *toolbarPanel {
	p := &toolbarPanel{
		parent:  u.window,
		current: u.current,
	}
	p.build(u.applyTheme)
	return p
}

func (p *toolbarPanel) build(applyThemeFunc func(fyne.Theme)) {
	themeSelect := widget.NewSelect(theme.EmbeddedThemes, func(string) {})
	themeSelect.SetSelectedIndex(0)
	themeApplyButton := widget.NewButton("Apply", func() {
		applyThemeFunc(theme.GetEmbeddedThemeFrom(themeSelect.Selected))
	})

	packageNameLabel := widget.NewLabel("Package name:")
	packageNameEntry := widget.NewEntry()
	packageNameEntry.SetText(p.current.PackageName())
	packageNameEntry.OnChanged = func(s string) { p.current.SetPackageName(s) }

	themeStructNameLabel := widget.NewLabel("Theme struct name:")
	themeStructNameEntry := widget.NewEntry()
	themeStructNameEntry.SetText(p.current.ThemeStructName())
	themeStructNameEntry.OnChanged = func(s string) { p.current.SetThemeStructName(s) }

	exportButton := widget.NewButton("Export theme", p.export)

	p.panel = fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			themeSelect,
			themeApplyButton,
		),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			packageNameLabel,
			packageNameEntry,
			themeStructNameLabel,
			themeStructNameEntry,
			exportButton,
		),
	)
}

func (p *toolbarPanel) export() {
	dstTheme, dstFont, err := theme.Generate(p.current)
	if err != nil {
		dialog.ShowError(err, p.parent)
		return
	}
	msg := fmt.Sprintf("Success to export file: %s", dstTheme)
	if dstFont != "" {
		msg += fmt.Sprintf(", %s", dstFont)
	}
	dialog.ShowInformation("Success", msg, p.parent)
}
