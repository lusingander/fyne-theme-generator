package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
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

const dummyVariant fyne.ThemeVariant = 100

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

	textSize           float32
	textFont           fyne.Resource
	textBoldFont       fyne.Resource
	textItalicFont     fyne.Resource
	textBoldItalicFont fyne.Resource
	textMonospaceFont  fyne.Resource

	padding            float32
	iconInlineSize     float32
	scrollBarSize      float32
	scrollBarSmallSize float32

	userSettingDefaultTheme fyne.Theme

	exportFontFile bool
	exportForV2    bool
}

func NewSetting() *Setting {
	userSettingDefaultTheme := fyne.CurrentApp().Settings().Theme()
	if userSettingDefaultTheme == nil {
		userSettingDefaultTheme = theme.DefaultTheme()
	}
	variant := fyne.CurrentApp().Settings().ThemeVariant()

	return &Setting{
		packageName:             defaultPackageName,
		themeStructName:         defaultThemeStructName,
		backgroundColor:         userSettingDefaultTheme.Color(theme.ColorNameBackground, variant),
		buttonColor:             userSettingDefaultTheme.Color(theme.ColorNameButton, variant),
		disabledButtonColor:     userSettingDefaultTheme.Color(theme.ColorNameDisabledButton, variant),
		textColor:               userSettingDefaultTheme.Color(theme.ColorNameForeground, variant),
		disabledTextColor:       userSettingDefaultTheme.Color(theme.ColorNameDisabled, variant),
		iconColor:               userSettingDefaultTheme.Color(theme.ColorNameForeground, variant),
		disabledIconColor:       userSettingDefaultTheme.Color(theme.ColorNameDisabled, variant),
		hyperlinkColor:          userSettingDefaultTheme.Color(theme.ColorNamePrimary, variant),
		placeHolderColor:        userSettingDefaultTheme.Color(theme.ColorNamePlaceHolder, variant),
		primaryColor:            userSettingDefaultTheme.Color(theme.ColorNamePrimary, variant),
		hoverColor:              userSettingDefaultTheme.Color(theme.ColorNameHover, variant),
		focusColor:              userSettingDefaultTheme.Color(theme.ColorNameFocus, variant),
		scrollBarColor:          userSettingDefaultTheme.Color(theme.ColorNameScrollBar, variant),
		shadowColor:             userSettingDefaultTheme.Color(theme.ColorNameShadow, variant),
		textSize:                userSettingDefaultTheme.Size(theme.SizeNameText),
		textFont:                userSettingDefaultTheme.Font(fyne.TextStyle{}),
		textBoldFont:            userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true}),
		textItalicFont:          userSettingDefaultTheme.Font(fyne.TextStyle{Italic: true}),
		textBoldItalicFont:      userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true, Italic: true}),
		textMonospaceFont:       userSettingDefaultTheme.Font(fyne.TextStyle{Monospace: true}),
		padding:                 userSettingDefaultTheme.Size(theme.SizeNamePadding),
		iconInlineSize:          userSettingDefaultTheme.Size(theme.SizeNameInlineIcon),
		scrollBarSize:           userSettingDefaultTheme.Size(theme.SizeNameScrollBar),
		scrollBarSmallSize:      userSettingDefaultTheme.Size(theme.SizeNameScrollBarSmall),
		userSettingDefaultTheme: userSettingDefaultTheme,
		exportFontFile:          false,
		exportForV2:             true,
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
func (s *Setting) TextSize() float32                 { return s.textSize }
func (s *Setting) TextFont() fyne.Resource           { return s.textFont }
func (s *Setting) TextBoldFont() fyne.Resource       { return s.textBoldFont }
func (s *Setting) TextItalicFont() fyne.Resource     { return s.textItalicFont }
func (s *Setting) TextBoldItalicFont() fyne.Resource { return s.textBoldItalicFont }
func (s *Setting) TextMonospaceFont() fyne.Resource  { return s.textMonospaceFont }
func (s *Setting) Padding() float32                  { return s.padding }
func (s *Setting) IconInlineSize() float32           { return s.iconInlineSize }
func (s *Setting) ScrollBarSize() float32            { return s.scrollBarSize }
func (s *Setting) ScrollBarSmallSize() float32       { return s.scrollBarSmallSize }

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
func (s *Setting) SetTextSize(size float32)                 { s.textSize = size }
func (s *Setting) SetTextFont(font fyne.Resource)           { s.textFont = font }
func (s *Setting) SetTextBoldFont(font fyne.Resource)       { s.textBoldFont = font }
func (s *Setting) SetTextItalicFont(font fyne.Resource)     { s.textItalicFont = font }
func (s *Setting) SetTextBoldItalicFont(font fyne.Resource) { s.textBoldItalicFont = font }
func (s *Setting) SetTextMonospaceFont(font fyne.Resource)  { s.textMonospaceFont = font }
func (s *Setting) SetPadding(pad float32)                   { s.padding = pad }
func (s *Setting) SetIconInlineSize(size float32)           { s.iconInlineSize = size }
func (s *Setting) SetScrollBarSize(size float32)            { s.scrollBarSize = size }
func (s *Setting) SetScrollBarSmallSize(size float32)       { s.scrollBarSmallSize = size }

func (s *Setting) ExportFontFile() bool     { return s.exportFontFile }
func (s *Setting) SetExportFontFile(b bool) { s.exportFontFile = b }
func (s *Setting) ExportForV2() bool        { return s.exportForV2 }
func (s *Setting) SetExportForV2(b bool)    { s.exportForV2 = b }

func (s *Setting) UpdateTheme(t fyne.Theme, v fyne.ThemeVariant) {
	s.SetBackgroundColor(t.Color(theme.ColorNameBackground, v))
	s.SetButtonColor(t.Color(theme.ColorNameButton, v))
	s.SetDisabledButtonColor(t.Color(theme.ColorNameDisabledButton, v))
	s.SetTextColor(t.Color(theme.ColorNameForeground, v))
	s.SetDisabledTextColor(t.Color(theme.ColorNameDisabled, v))
	s.SetIconColor(t.Color(theme.ColorNameForeground, v))
	s.SetDisabledIconColor(t.Color(theme.ColorNameDisabled, v))
	s.SetHyperlinkColor(t.Color(theme.ColorNamePrimary, v))
	s.SetPlaceHolderColor(t.Color(theme.ColorNamePlaceHolder, v))
	s.SetPrimaryColor(t.Color(theme.ColorNamePrimary, v))
	s.SetHoverColor(t.Color(theme.ColorNameHover, v))
	s.SetFocusColor(t.Color(theme.ColorNameFocus, v))
	s.SetScrollBarColor(t.Color(theme.ColorNameScrollBar, v))
	s.SetShadowColor(t.Color(theme.ColorNameShadow, v))
	s.SetTextSize(t.Size(theme.SizeNameText))
	s.SetTextFont(t.Font(fyne.TextStyle{}))
	s.SetTextBoldFont(t.Font(fyne.TextStyle{Bold: true}))
	s.SetTextItalicFont(t.Font(fyne.TextStyle{Italic: true}))
	s.SetTextBoldItalicFont(t.Font(fyne.TextStyle{Bold: true, Italic: true}))
	s.SetTextMonospaceFont(t.Font(fyne.TextStyle{Monospace: true}))
	s.SetPadding(t.Size(theme.SizeNamePadding))
	s.SetIconInlineSize(t.Size(theme.SizeNameInlineIcon))
	s.SetScrollBarSize(t.Size(theme.SizeNameScrollBar))
	s.SetScrollBarSmallSize(t.Size(theme.SizeNameScrollBarSmall))
}

func (s *Setting) isSetTextFont() bool {
	return s.userSettingDefaultTheme.Font(fyne.TextStyle{}).Name() != s.textFont.Name()
}

func (s *Setting) isSetTextBoldFont() bool {
	return s.userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true}).Name() != s.textBoldFont.Name()
}

func (s *Setting) isSetTextItalicFont() bool {
	return s.userSettingDefaultTheme.Font(fyne.TextStyle{Italic: true}).Name() != s.textItalicFont.Name()
}

func (s *Setting) isSetTextBoldItalicFont() bool {
	return s.userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true, Italic: true}).Name() != s.textBoldItalicFont.Name()
}

func (s *Setting) isSetTextMonospaceFont() bool {
	return s.userSettingDefaultTheme.Font(fyne.TextStyle{Monospace: true}).Name() != s.textMonospaceFont.Name()
}

func (s *Setting) isSetOneOfFonts() bool {
	return s.isSetTextFont() || s.isSetTextBoldFont() || s.isSetTextItalicFont() || s.isSetTextBoldItalicFont() || s.isSetTextMonospaceFont()
}

func (s *Setting) needToGenerateFont() bool {
	return s.exportFontFile && s.isSetOneOfFonts()
}

func (s *Setting) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return s.backgroundColor
	// case theme.ColorNameInputBackground:
	// 	return mil
	case theme.ColorNameForeground:
		return s.textColor
	case theme.ColorNameButton:
		return s.buttonColor
	case theme.ColorNameDisabled:
		return s.disabledTextColor
	case theme.ColorNameDisabledButton:
		return s.disabledButtonColor
	case theme.ColorNamePlaceHolder:
		return s.placeHolderColor
	case theme.ColorNameScrollBar:
		return s.scrollBarColor
	case theme.ColorNamePrimary:
		return s.primaryColor
	case theme.ColorNameHover:
		return s.hoverColor
	case theme.ColorNameFocus:
		return s.focusColor
	case theme.ColorNameShadow:
		return s.shadowColor
	// case theme.ColorNameError:
	// 	return nil
	// case theme.ColorNamePressed:
	// 	return nil
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (s *Setting) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return s.textMonospaceFont
	}
	if style.Bold {
		if style.Italic {
			return s.textBoldFont
		}
		return s.textBoldItalicFont
	}
	if style.Italic {
		return s.textItalicFont
	}
	return s.textFont
}

func (s *Setting) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (s *Setting) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNamePadding:
		return s.padding
	case theme.SizeNameInlineIcon:
		return s.iconInlineSize
	case theme.SizeNameScrollBar:
		return s.scrollBarSize
	case theme.SizeNameScrollBarSmall:
		return s.scrollBarSmallSize
	case theme.SizeNameText:
		return s.textSize
	default:
		return theme.DefaultTheme().Size(n)
	}
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
