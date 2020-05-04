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

func export(base fyne.Window) {
	dst, err := generate(currentThemeSetting)
	if err != nil {
		dialog.ShowError(err, base)
		return
	}
	msg := fmt.Sprintf("Success to export file: %s", dst)
	dialog.ShowInformation("Success", msg, base)
}

func run(args []string) error {
	a := app.New()
	ts := newThemeSetting()
	currentThemeSetting = ts
	a.Settings().SetTheme(ts)
	w := a.NewWindow("Fyne theme generator")
	mainWindow = w
	confs := configures(ts)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				fyne.NewContainerWithLayout(
					layout.NewGridLayoutWithColumns(3),
					confs[:len(confs)/2]...,
				),
				fyne.NewContainerWithLayout(
					layout.NewGridLayoutWithColumns(3),
					confs[len(confs)/2:]...,
				),
			),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				widget.NewButton("Export theme", func() { export(w) }),
			),
		),
	)
	w.ShowAndRun()
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
