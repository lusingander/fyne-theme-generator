package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

func (u *ui) export() {
	dst, err := theme.Generate(u.current)
	if err != nil {
		dialog.ShowError(err, u.window)
		return
	}
	msg := fmt.Sprintf("Success to export file: %s", dst)
	dialog.ShowInformation("Success", msg, u.window)
}

func (u *ui) toolbar() fyne.CanvasObject {
	packageNameLabel := widget.NewLabel("Package name:")
	packageNameEntry := widget.NewEntry()
	packageNameEntry.SetText(u.current.PackageName())
	packageNameEntry.OnChanged = func(s string) { u.current.SetPackageName(s) }

	themeStructNameLabel := widget.NewLabel("Theme struct name:")
	themeStructNameEntry := widget.NewEntry()
	themeStructNameEntry.SetText(u.current.ThemeStructName())
	themeStructNameEntry.OnChanged = func(s string) { u.current.SetThemeStructName(s) }

	exportButton := widget.NewButton("Export theme", u.export)

	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		packageNameLabel,
		packageNameEntry,
		themeStructNameLabel,
		themeStructNameEntry,
		exportButton,
	)
}
