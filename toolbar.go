package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func export() {
	dst, err := generate(currentThemeSetting)
	if err != nil {
		dialog.ShowError(err, mainWindow)
		return
	}
	msg := fmt.Sprintf("Success to export file: %s", dst)
	dialog.ShowInformation("Success", msg, mainWindow)
}

func toolbar() fyne.CanvasObject {
	packageNameLabel := widget.NewLabel("Package name:")
	packageNameEntry := widget.NewEntry()
	packageNameEntry.SetText(currentThemeSetting.PackageName())
	packageNameEntry.OnChanged = func(s string) { currentThemeSetting.SetPackageName(s) }

	themeStructNameLabel := widget.NewLabel("Theme struct name:")
	themeStructNameEntry := widget.NewEntry()
	themeStructNameEntry.SetText(currentThemeSetting.ThemeStructName())
	themeStructNameEntry.OnChanged = func(s string) { currentThemeSetting.SetThemeStructName(s) }

	exportButton := widget.NewButton("Export theme", export)

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
