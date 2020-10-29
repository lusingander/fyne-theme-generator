package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

type ui struct {
	window  fyne.Window
	current *theme.Setting
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
	u.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				u.config(),
				u.toolbar(),
			),
			u.widgets(),
		),
	)
}
