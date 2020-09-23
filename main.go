package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var (
	currentThemeSetting *themeSetting
	mainWindow          fyne.Window
)

func reflesh() {
	fyne.CurrentApp().Settings().SetTheme(currentThemeSetting)
}

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

func run(args []string) error {
	a := app.New()
	currentThemeSetting = newThemeSetting()
	a.Settings().SetTheme(currentThemeSetting)
	mainWindow = a.NewWindow("Fyne theme generator")
	confs := configures(currentThemeSetting)
	mainWindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				fyne.NewContainerWithLayout(
					layout.NewHBoxLayout(),
					fyne.NewContainerWithLayout(
						layout.NewGridLayoutWithColumns(2),
						confs[:len(confs)/2]...,
					),
					fyne.NewContainerWithLayout(
						layout.NewGridLayoutWithColumns(2),
						confs[len(confs)/2:]...,
					),
				),
				toolbar(),
			),
			widgets(),
		),
	)
	mainWindow.ShowAndRun()
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
