package theme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

const (
	defaultPackageName     = "main"
	defaultThemeStructName = "myTheme"
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

func GetEmbeddedThemeFrom(name string) fyne.Theme {
	switch name {
	case fyneDarkThemeName:
		return theme.DarkTheme()
	case fyneLightThemeName:
		return theme.LightTheme()
	case fyneOldDarkThemeName:
		return &fyneOldDarkTheme{}
	case fyneOldLightThemeName:
		return &fyneOldLightTheme{}
	}
	return theme.DarkTheme()
}

type Setting struct {
	packageName     string
	themeStructName string

	backgroundColor     color.Color
	buttonColor         color.Color
	disabledButtonColor color.Color
	textColor           color.Color
	disabledTextColor   color.Color
	iconColor           color.Color
	disabledIconColor   color.Color
	hyperlinkColor      color.Color
	placeHolderColor    color.Color
	primaryColor        color.Color
	hoverColor          color.Color
	focusColor          color.Color
	scrollBarColor      color.Color
	shadowColor         color.Color

	textSize           int
	textFont           fyne.Resource
	textBoldFont       fyne.Resource
	textItalicFont     fyne.Resource
	textBoldItalicFont fyne.Resource
	textMonospaceFont  fyne.Resource

	padding            int
	iconInlineSize     int
	scrollBarSize      int
	scrollBarSmallSize int

	userSettingDefaultTheme fyne.Theme
}

func NewSetting() *Setting {
	userSettingDefaultTheme := fyne.CurrentApp().Settings().Theme()
	if userSettingDefaultTheme == nil {
		userSettingDefaultTheme = theme.DarkTheme()
	}

	return &Setting{
		packageName:             defaultPackageName,
		themeStructName:         defaultThemeStructName,
		backgroundColor:         userSettingDefaultTheme.BackgroundColor(),
		buttonColor:             userSettingDefaultTheme.ButtonColor(),
		disabledButtonColor:     userSettingDefaultTheme.DisabledButtonColor(),
		textColor:               userSettingDefaultTheme.TextColor(),
		disabledTextColor:       userSettingDefaultTheme.DisabledTextColor(),
		iconColor:               userSettingDefaultTheme.IconColor(),
		disabledIconColor:       userSettingDefaultTheme.DisabledIconColor(),
		hyperlinkColor:          userSettingDefaultTheme.HyperlinkColor(),
		placeHolderColor:        userSettingDefaultTheme.PlaceHolderColor(),
		primaryColor:            userSettingDefaultTheme.PrimaryColor(),
		hoverColor:              userSettingDefaultTheme.HoverColor(),
		focusColor:              userSettingDefaultTheme.FocusColor(),
		scrollBarColor:          userSettingDefaultTheme.ScrollBarColor(),
		shadowColor:             userSettingDefaultTheme.ShadowColor(),
		textSize:                userSettingDefaultTheme.TextSize(),
		textFont:                userSettingDefaultTheme.TextFont(),
		textBoldFont:            userSettingDefaultTheme.TextBoldFont(),
		textItalicFont:          userSettingDefaultTheme.TextItalicFont(),
		textBoldItalicFont:      userSettingDefaultTheme.TextBoldItalicFont(),
		textMonospaceFont:       userSettingDefaultTheme.TextMonospaceFont(),
		padding:                 userSettingDefaultTheme.Padding(),
		iconInlineSize:          userSettingDefaultTheme.IconInlineSize(),
		scrollBarSize:           userSettingDefaultTheme.ScrollBarSize(),
		scrollBarSmallSize:      userSettingDefaultTheme.ScrollBarSmallSize(),
		userSettingDefaultTheme: userSettingDefaultTheme,
	}
}

func (s *Setting) PackageName() string               { return s.packageName }
func (s *Setting) ThemeStructName() string           { return s.themeStructName }
func (s *Setting) BackgroundColor() color.Color      { return s.backgroundColor }
func (s *Setting) ButtonColor() color.Color          { return s.buttonColor }
func (s *Setting) DisabledButtonColor() color.Color  { return s.disabledButtonColor }
func (s *Setting) TextColor() color.Color            { return s.textColor }
func (s *Setting) DisabledTextColor() color.Color    { return s.disabledTextColor }
func (s *Setting) IconColor() color.Color            { return s.iconColor }
func (s *Setting) DisabledIconColor() color.Color    { return s.disabledIconColor }
func (s *Setting) HyperlinkColor() color.Color       { return s.hyperlinkColor }
func (s *Setting) PlaceHolderColor() color.Color     { return s.placeHolderColor }
func (s *Setting) PrimaryColor() color.Color         { return s.primaryColor }
func (s *Setting) HoverColor() color.Color           { return s.hoverColor }
func (s *Setting) FocusColor() color.Color           { return s.focusColor }
func (s *Setting) ScrollBarColor() color.Color       { return s.scrollBarColor }
func (s *Setting) ShadowColor() color.Color          { return s.shadowColor }
func (s *Setting) TextSize() int                     { return s.textSize }
func (s *Setting) TextFont() fyne.Resource           { return s.textFont }
func (s *Setting) TextBoldFont() fyne.Resource       { return s.textBoldFont }
func (s *Setting) TextItalicFont() fyne.Resource     { return s.textItalicFont }
func (s *Setting) TextBoldItalicFont() fyne.Resource { return s.textBoldItalicFont }
func (s *Setting) TextMonospaceFont() fyne.Resource  { return s.textMonospaceFont }
func (s *Setting) Padding() int                      { return s.padding }
func (s *Setting) IconInlineSize() int               { return s.iconInlineSize }
func (s *Setting) ScrollBarSize() int                { return s.scrollBarSize }
func (s *Setting) ScrollBarSmallSize() int           { return s.scrollBarSmallSize }

func (s *Setting) SetPackageName(name string)               { s.packageName = name }
func (s *Setting) SetThemeStructName(name string)           { s.themeStructName = name }
func (s *Setting) SetBackgroundColor(c color.Color)         { s.backgroundColor = c }
func (s *Setting) SetButtonColor(c color.Color)             { s.buttonColor = c }
func (s *Setting) SetDisabledButtonColor(c color.Color)     { s.disabledButtonColor = c }
func (s *Setting) SetTextColor(c color.Color)               { s.textColor = c }
func (s *Setting) SetDisabledTextColor(c color.Color)       { s.disabledTextColor = c }
func (s *Setting) SetIconColor(c color.Color)               { s.iconColor = c }
func (s *Setting) SetDisabledIconColor(c color.Color)       { s.disabledIconColor = c }
func (s *Setting) SetHyperlinkColor(c color.Color)          { s.hyperlinkColor = c }
func (s *Setting) SetPlaceHolderColor(c color.Color)        { s.placeHolderColor = c }
func (s *Setting) SetPrimaryColor(c color.Color)            { s.primaryColor = c }
func (s *Setting) SetHoverColor(c color.Color)              { s.hoverColor = c }
func (s *Setting) SetFocusColor(c color.Color)              { s.focusColor = c }
func (s *Setting) SetScrollBarColor(c color.Color)          { s.scrollBarColor = c }
func (s *Setting) SetShadowColor(c color.Color)             { s.shadowColor = c }
func (s *Setting) SetTextSize(size int)                     { s.textSize = size }
func (s *Setting) SetTextFont(font fyne.Resource)           { s.textFont = font }
func (s *Setting) SetTextBoldFont(font fyne.Resource)       { s.textBoldFont = font }
func (s *Setting) SetTextItalicFont(font fyne.Resource)     { s.textItalicFont = font }
func (s *Setting) SetTextBoldItalicFont(font fyne.Resource) { s.textBoldItalicFont = font }
func (s *Setting) SetTextMonospaceFont(font fyne.Resource)  { s.textMonospaceFont = font }
func (s *Setting) SetPadding(pad int)                       { s.padding = pad }
func (s *Setting) SetIconInlineSize(size int)               { s.iconInlineSize = size }
func (s *Setting) SetScrollBarSize(size int)                { s.scrollBarSize = size }
func (s *Setting) SetScrollBarSmallSize(size int)           { s.scrollBarSmallSize = size }

func (s *Setting) UpdateTheme(t fyne.Theme) {
	s.SetBackgroundColor(t.BackgroundColor())
	s.SetButtonColor(t.ButtonColor())
	s.SetDisabledButtonColor(t.DisabledButtonColor())
	s.SetTextColor(t.TextColor())
	s.SetDisabledTextColor(t.DisabledTextColor())
	s.SetIconColor(t.IconColor())
	s.SetDisabledIconColor(t.DisabledIconColor())
	s.SetHyperlinkColor(t.HyperlinkColor())
	s.SetPlaceHolderColor(t.PlaceHolderColor())
	s.SetPrimaryColor(t.PrimaryColor())
	s.SetHoverColor(t.HoverColor())
	s.SetFocusColor(t.FocusColor())
	s.SetScrollBarColor(t.ScrollBarColor())
	s.SetShadowColor(t.ShadowColor())
	s.SetTextSize(t.TextSize())
	s.SetTextFont(t.TextFont())
	s.SetTextBoldFont(t.TextBoldFont())
	s.SetTextItalicFont(t.TextItalicFont())
	s.SetTextBoldItalicFont(t.TextBoldItalicFont())
	s.SetTextMonospaceFont(t.TextMonospaceFont())
	s.SetPadding(t.Padding())
	s.SetIconInlineSize(t.IconInlineSize())
	s.SetScrollBarSize(t.ScrollBarSize())
	s.SetScrollBarSmallSize(t.ScrollBarSmallSize())
}

func (s *Setting) isSetTextFont() bool {
	return s.userSettingDefaultTheme.TextFont().Name() != s.textFont.Name()
}

func (s *Setting) isSetTextBoldFont() bool {
	return s.userSettingDefaultTheme.TextBoldFont().Name() != s.textBoldFont.Name()
}

func (s *Setting) isSetTextItalicFont() bool {
	return s.userSettingDefaultTheme.TextItalicFont().Name() != s.textItalicFont.Name()
}

func (s *Setting) isSetTextBoldItalicFont() bool {
	return s.userSettingDefaultTheme.TextBoldItalicFont().Name() != s.textBoldItalicFont.Name()
}

func (s *Setting) isSetTextMonospaceFont() bool {
	return s.userSettingDefaultTheme.TextMonospaceFont().Name() != s.textMonospaceFont.Name()
}

func (s *Setting) needToGenerateFont() bool {
	return s.isSetTextFont() ||
		s.isSetTextBoldFont() ||
		s.isSetTextItalicFont() ||
		s.isSetTextBoldItalicFont() ||
		s.isSetTextMonospaceFont()
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
	return fyne.CurrentApp().Settings().Theme().TextFont()
}

func (fyneOldDarkTheme) TextBoldFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextBoldFont()
}

func (fyneOldDarkTheme) TextItalicFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextItalicFont()
}

func (fyneOldDarkTheme) TextBoldItalicFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextBoldItalicFont()
}

func (fyneOldDarkTheme) TextMonospaceFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextMonospaceFont()
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
	return fyne.CurrentApp().Settings().Theme().TextFont()
}

func (fyneOldLightTheme) TextBoldFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextBoldFont()
}

func (fyneOldLightTheme) TextItalicFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextItalicFont()
}

func (fyneOldLightTheme) TextBoldItalicFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextBoldItalicFont()
}

func (fyneOldLightTheme) TextMonospaceFont() fyne.Resource {
	return fyne.CurrentApp().Settings().Theme().TextMonospaceFont()
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
