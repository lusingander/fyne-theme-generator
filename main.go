package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/lusingander/fyne-theme-generator/internal/ui"
)

const (
	appTitle = "Fyne theme generator"
)

func run(args []string) error {
	a := app.New()
	w := a.NewWindow(appTitle)
	return ui.Start(w)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
