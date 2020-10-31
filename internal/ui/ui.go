package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

type ui struct {
	window  fyne.Window
	current *theme.Setting

	*configPanel
	*toolbarPanel
	*widgetsPanel
}

func New(w fyne.Window) *ui {
	return &ui{
		window:  w,
		current: theme.NewSetting(),
	}
}

func (u *ui) Reflesh() {
	fyne.CurrentApp().Settings().SetTheme(u.current)
}

func (u *ui) Build() {
	u.configPanel = u.newConfigPanel()
	u.toolbarPanel = u.newToolbarPanel()
	u.widgetsPanel = u.newWidgetsPanel()
	u.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				u.configPanel.panel,
				u.toolbarPanel.panel,
			),
			u.widgetsPanel.panel,
		),
	)
}

func (u *ui) applyTheme(t fyne.Theme) {
	u.current.UpdateTheme(t)
	u.configPanel.applyCurrentTheme()
	u.Reflesh()
}
