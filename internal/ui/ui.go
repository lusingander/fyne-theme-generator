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

	widgetsPanelWindow fyne.Window
}

// Start app
func Start(w fyne.Window) {
	new(w).build()
}

func new(w fyne.Window) *ui {
	return &ui{
		window:  w,
		current: theme.NewSetting(),
	}
}

func (u *ui) reflesh() {
	fyne.CurrentApp().Settings().SetTheme(u.current)
}

func (u *ui) build() {
	u.configPanel = u.newConfigPanel()
	u.toolbarPanel = u.newToolbarPanel()
	u.widgetsPanel = u.newWidgetsPanel()
	u.setContent(true)
}

func (u *ui) setContent(showWidgets bool) {
	container := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			u.configPanel.panel,
			u.toolbarPanel.panel,
		),
	)
	if showWidgets {
		container.Add(u.widgetsPanel.panel)
	}
	u.window.SetContent(container)
}

func (u *ui) applyTheme(t fyne.Theme) {
	u.current.UpdateTheme(t)
	u.configPanel.applyCurrentTheme()
	u.reflesh()
}

func (u *ui) connectWidgetPanel() {
	u.setContent(true)
	u.widgetsPanel.connect(u.disconnectWidgetPanel)
	u.widgetsPanelWindow.Close()
	u.reflesh()
}

func (u *ui) disconnectWidgetPanel() {
	u.setContent(false)
	u.widgetsPanel.disconnect(u.connectWidgetPanel)
	u.widgetsPanelWindow = fyne.CurrentApp().NewWindow("widgets")
	u.widgetsPanelWindow.SetContent(u.widgetsPanel.panel)
	u.widgetsPanelWindow.Show()
	u.window.Resize(u.window.Content().MinSize())
	u.reflesh()
}
