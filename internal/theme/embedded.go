package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	fyneDarkThemeName     = "Fyne Dark"
	fyneLightThemeName    = "Fyne Light"
	fyneOldDarkThemeName  = "Fyne Dark (old)"
	fyneOldLightThemeName = "Fyne Light (old)"
)

var EmbeddedThemes = []string{
	fyneDarkThemeName,
	fyneLightThemeName,
	fyneOldDarkThemeName,
	fyneOldLightThemeName,
}

func GetEmbeddedThemeFrom(name string) (fyne.Theme, fyne.ThemeVariant) {
	switch name {
	case fyneDarkThemeName:
		return theme.DarkTheme(), theme.VariantDark
	case fyneLightThemeName:
		return theme.LightTheme(), theme.VariantLight
	case fyneOldDarkThemeName:
		return theme.FromLegacy(&fyneOldDarkTheme{}), dummyVariant
	case fyneOldLightThemeName:
		return theme.FromLegacy(&fyneOldLightTheme{}), dummyVariant
	}
	return theme.DefaultTheme(), dummyVariant
}

type fyneOldDarkTheme struct{}

func (fyneOldDarkTheme) BackgroundColor() color.Color {
	return color.NRGBA{0x42, 0x42, 0x42, 0xff}
}

func (fyneOldDarkTheme) ButtonColor() color.Color {
	return color.NRGBA{0x21, 0x21, 0x21, 0xff}
}

func (fyneOldDarkTheme) DisabledButtonColor() color.Color {
	return color.NRGBA{0x31, 0x31, 0x31, 0xff}
}

func (fyneOldDarkTheme) IconColor() color.Color {
	return color.NRGBA{0xff, 0xff, 0xff, 0xff}
}

func (fyneOldDarkTheme) DisabledIconColor() color.Color {
	return color.NRGBA{0x60, 0x60, 0x60, 0xff}
}

func (fyneOldDarkTheme) HyperlinkColor() color.Color {
	return color.NRGBA{0x99, 0x99, 0xff, 0xff}
}

func (fyneOldDarkTheme) TextColor() color.Color {
	return color.NRGBA{0xff, 0xff, 0xff, 0xff}
}

func (fyneOldDarkTheme) DisabledTextColor() color.Color {
	return color.NRGBA{0x60, 0x60, 0x60, 0xff}
}

func (fyneOldDarkTheme) HoverColor() color.Color {
	return color.NRGBA{0x31, 0x31, 0x31, 0xff}
}

func (fyneOldDarkTheme) PlaceHolderColor() color.Color {
	return color.NRGBA{0xb2, 0xb2, 0xb2, 0xff}
}

func (fyneOldDarkTheme) PrimaryColor() color.Color {
	return color.NRGBA{0x1a, 0x23, 0x7e, 0xff}
}

func (fyneOldDarkTheme) FocusColor() color.Color {
	return color.NRGBA{0x1a, 0x23, 0x7e, 0xff}
}

func (fyneOldDarkTheme) ScrollBarColor() color.Color {
	return color.NRGBA{0x0, 0x0, 0x0, 0x99}
}

func (fyneOldDarkTheme) ShadowColor() color.Color {
	return color.NRGBA{0x0, 0x0, 0x0, 0x66}
}

func (fyneOldDarkTheme) TextSize() int {
	return 14
}

func (fyneOldDarkTheme) TextFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{})
}

func (fyneOldDarkTheme) TextBoldFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Bold: true})
}

func (fyneOldDarkTheme) TextItalicFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Italic: true})
}

func (fyneOldDarkTheme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Bold: true, Italic: true})
}

func (fyneOldDarkTheme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Monospace: true})
}

func (fyneOldDarkTheme) Padding() int {
	return 4
}

func (fyneOldDarkTheme) IconInlineSize() int {
	return 20
}

func (fyneOldDarkTheme) ScrollBarSize() int {
	return 16
}

func (fyneOldDarkTheme) ScrollBarSmallSize() int {
	return 3
}

type fyneOldLightTheme struct{}

func (fyneOldLightTheme) BackgroundColor() color.Color {
	return color.NRGBA{0xf5, 0xf5, 0xf5, 0xff}
}

func (fyneOldLightTheme) ButtonColor() color.Color {
	return color.NRGBA{0xd9, 0xd9, 0xd9, 0xff}
}

func (fyneOldLightTheme) DisabledButtonColor() color.Color {
	return color.NRGBA{0xe7, 0xe7, 0xe7, 0xff}
}

func (fyneOldLightTheme) IconColor() color.Color {
	return color.NRGBA{0x21, 0x21, 0x21, 0xff}
}

func (fyneOldLightTheme) DisabledIconColor() color.Color {
	return color.NRGBA{0x80, 0x80, 0x80, 0xff}
}

func (fyneOldLightTheme) HyperlinkColor() color.Color {
	return color.NRGBA{0x0, 0x0, 0xd9, 0xff}
}

func (fyneOldLightTheme) TextColor() color.Color {
	return color.NRGBA{0x21, 0x21, 0x21, 0xff}
}

func (fyneOldLightTheme) DisabledTextColor() color.Color {
	return color.NRGBA{0x80, 0x80, 0x80, 0xff}
}

func (fyneOldLightTheme) HoverColor() color.Color {
	return color.NRGBA{0xe7, 0xe7, 0xe7, 0xff}
}

func (fyneOldLightTheme) PlaceHolderColor() color.Color {
	return color.NRGBA{0x88, 0x88, 0x88, 0xff}
}

func (fyneOldLightTheme) PrimaryColor() color.Color {
	return color.NRGBA{0x9f, 0xa8, 0xda, 0xff}
}

func (fyneOldLightTheme) FocusColor() color.Color {
	return color.NRGBA{0x9f, 0xa8, 0xda, 0xff}
}

func (fyneOldLightTheme) ScrollBarColor() color.Color {
	return color.NRGBA{0x0, 0x0, 0x0, 0x99}
}

func (fyneOldLightTheme) ShadowColor() color.Color {
	return color.NRGBA{0x0, 0x0, 0x0, 0x33}
}

func (fyneOldLightTheme) TextSize() int {
	return 14
}

func (fyneOldLightTheme) TextFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{})
}

func (fyneOldLightTheme) TextBoldFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Bold: true})
}

func (fyneOldLightTheme) TextItalicFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Italic: true})
}

func (fyneOldLightTheme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Bold: true, Italic: true})
}

func (fyneOldLightTheme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Monospace: true})
}

func (fyneOldLightTheme) Padding() int {
	return 4
}

func (fyneOldLightTheme) IconInlineSize() int {
	return 20
}

func (fyneOldLightTheme) ScrollBarSize() int {
	return 16
}

func (fyneOldLightTheme) ScrollBarSmallSize() int {
	return 3
}
