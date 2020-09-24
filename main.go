package main

import (
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
)

const (
	appTitle = "Fyne theme generator"
)

var (
	currentThemeSetting *themeSetting
	mainWindow          fyne.Window
)

func reflesh() {
	fyne.CurrentApp().Settings().SetTheme(currentThemeSetting)
}

func run(args []string) error {
	a := app.New()
	currentThemeSetting = newThemeSetting()
	a.Settings().SetTheme(currentThemeSetting)
	mainWindow = a.NewWindow(appTitle)
	mainWindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				config(),
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
