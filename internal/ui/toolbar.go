package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	ft "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	p.build(u.applyTheme, u.refresh)
	return p
}

func (p *toolbarPanel) build(applyThemeFunc func(fyne.Theme, fyne.ThemeVariant), refreshFunc func()) {
	themeSelect := widget.NewSelect(theme.EmbeddedThemes, func(string) {})
	themeSelect.SetSelectedIndex(0)
	themeApplyButton := widget.NewButton("Apply", func() {
		applyThemeFunc(theme.GetEmbeddedThemeFrom(themeSelect.Selected))
	})

	refreshButton := widget.NewButtonWithIcon("", ft.ViewRefreshIcon(), refreshFunc)

	packageNameLabel := widget.NewLabel("Package name:")
	packageNameEntry := &widget.Entry{Wrapping: fyne.TextWrapOff}
	packageNameEntry.SetText(p.current.PackageName())
	packageNameEntry.OnChanged = func(s string) { p.current.SetPackageName(s) }

	themeStructNameLabel := widget.NewLabel("Theme struct name:")
	themeStructNameEntry := &widget.Entry{Wrapping: fyne.TextWrapOff}
	themeStructNameEntry.SetText(p.current.ThemeStructName())
	themeStructNameEntry.OnChanged = func(s string) { p.current.SetThemeStructName(s) }

	fyneVerLabel := widget.NewLabel("Target Fyne version:")
	fyneVerSelect := widget.NewSelect([]string{"v1", "v2"}, func(s string) {
		p.current.SetExportForV2(s == "v2")
	})
	fyneVerSelect.SetSelectedIndex(1)

	exportFontCheck := widget.NewCheck("Generate font file", p.current.SetExportFontFile)

	exportButton := widget.NewButton("Export theme", p.export)

	p.panel = container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewHBoxLayout(),
			themeSelect,
			themeApplyButton,
			layout.NewSpacer(),
			packageNameLabel,
			packageNameEntry,
			themeStructNameLabel,
			themeStructNameEntry,
		),
		container.New(
			layout.NewHBoxLayout(),
			refreshButton,
			layout.NewSpacer(),
			fyneVerLabel,
			fyneVerSelect,
			exportFontCheck,
			exportButton,
		),
	)
}

func (p *toolbarPanel) export() {
	prog := dialog.NewProgressInfinite("Export", "exporting...", p.parent)
	prog.Show()
	dstTheme, dstFont, err := theme.Generate(p.current)
	prog.Hide()
	if err != nil {
		dialog.ShowError(err, p.parent)
		return
	}
	msg := fmt.Sprintf("Success to export file: %s", dstTheme)
	if dstFont != "" {
		msg += fmt.Sprintf(", %s", dstFont)
	} else {
		if p.current.ExportFontFile() {
			msg += "\n\nFont file is not exported because the font has not been changed..."
		}
	}
	dialog.ShowInformation("Success", msg, p.parent)
}
