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

	backgroundColor      color.Color
	buttonColor          color.Color
	disabledButtonColor  color.Color
	disabledColor        color.Color
	errorColor           color.Color
	focusColor           color.Color
	foregroundColor      color.Color
	hoverColor           color.Color
	inputBackgroundColor color.Color
	placeHolderColor     color.Color
	pressedColor         color.Color
	primaryColor         color.Color
	scrollBarColor       color.Color
	shadowColor          color.Color

	captionTextSize        float32
	inlineIconSize         float32
	paddingSize            float32
	scrollBarSize          float32
	scrollBarSmallSize     float32
	separatorThicknessSize float32
	textSize               float32
	inputBorderSize        float32

	textFont           fyne.Resource
	textBoldFont       fyne.Resource
	textItalicFont     fyne.Resource
	textBoldItalicFont fyne.Resource
	textMonospaceFont  fyne.Resource

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
		packageName:     defaultPackageName,
		themeStructName: defaultThemeStructName,

		backgroundColor:      userSettingDefaultTheme.Color(theme.ColorNameBackground, variant),
		buttonColor:          userSettingDefaultTheme.Color(theme.ColorNameButton, variant),
		disabledButtonColor:  userSettingDefaultTheme.Color(theme.ColorNameDisabledButton, variant),
		disabledColor:        userSettingDefaultTheme.Color(theme.ColorNameDisabled, variant),
		errorColor:           userSettingDefaultTheme.Color(theme.ColorNameError, variant),
		focusColor:           userSettingDefaultTheme.Color(theme.ColorNameFocus, variant),
		foregroundColor:      userSettingDefaultTheme.Color(theme.ColorNameForeground, variant),
		hoverColor:           userSettingDefaultTheme.Color(theme.ColorNameHover, variant),
		inputBackgroundColor: userSettingDefaultTheme.Color(theme.ColorNameInputBackground, variant),
		placeHolderColor:     userSettingDefaultTheme.Color(theme.ColorNamePlaceHolder, variant),
		pressedColor:         userSettingDefaultTheme.Color(theme.ColorNamePressed, variant),
		primaryColor:         userSettingDefaultTheme.Color(theme.ColorNamePrimary, variant),
		scrollBarColor:       userSettingDefaultTheme.Color(theme.ColorNameScrollBar, variant),
		shadowColor:          userSettingDefaultTheme.Color(theme.ColorNameShadow, variant),

		captionTextSize:        userSettingDefaultTheme.Size(theme.SizeNameCaptionText),
		inlineIconSize:         userSettingDefaultTheme.Size(theme.SizeNameInlineIcon),
		paddingSize:            userSettingDefaultTheme.Size(theme.SizeNamePadding),
		scrollBarSize:          userSettingDefaultTheme.Size(theme.SizeNameScrollBar),
		scrollBarSmallSize:     userSettingDefaultTheme.Size(theme.SizeNameScrollBarSmall),
		separatorThicknessSize: userSettingDefaultTheme.Size(theme.SizeNameSeparatorThickness),
		textSize:               userSettingDefaultTheme.Size(theme.SizeNameText),
		inputBorderSize:        userSettingDefaultTheme.Size(theme.SizeNameInputBorder),

		textFont:           userSettingDefaultTheme.Font(fyne.TextStyle{}),
		textBoldFont:       userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true}),
		textItalicFont:     userSettingDefaultTheme.Font(fyne.TextStyle{Italic: true}),
		textBoldItalicFont: userSettingDefaultTheme.Font(fyne.TextStyle{Bold: true, Italic: true}),
		textMonospaceFont:  userSettingDefaultTheme.Font(fyne.TextStyle{Monospace: true}),

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
func (s *Setting) DisabledColor() color.Color        { return s.disabledColor }
func (s *Setting) ErrorColor() color.Color           { return s.errorColor }
func (s *Setting) FocusColor() color.Color           { return s.focusColor }
func (s *Setting) ForegroundColor() color.Color      { return s.foregroundColor }
func (s *Setting) HoverColor() color.Color           { return s.hoverColor }
func (s *Setting) InputBackgroundColor() color.Color { return s.inputBackgroundColor }
func (s *Setting) PlaceHolderColor() color.Color     { return s.placeHolderColor }
func (s *Setting) PressedColor() color.Color         { return s.pressedColor }
func (s *Setting) PrimaryColor() color.Color         { return s.primaryColor }
func (s *Setting) ScrollBarColor() color.Color       { return s.scrollBarColor }
func (s *Setting) ShadowColor() color.Color          { return s.shadowColor }
func (s *Setting) CaptionTextSize() float32          { return s.captionTextSize }
func (s *Setting) InlineIconSize() float32           { return s.inlineIconSize }
func (s *Setting) PaddingSize() float32              { return s.paddingSize }
func (s *Setting) ScrollBarSize() float32            { return s.scrollBarSize }
func (s *Setting) ScrollBarSmallSize() float32       { return s.scrollBarSmallSize }
func (s *Setting) SeparatorThicknessSize() float32   { return s.separatorThicknessSize }
func (s *Setting) TextSize() float32                 { return s.textSize }
func (s *Setting) InputBorderSize() float32          { return s.inputBorderSize }
func (s *Setting) TextFont() fyne.Resource           { return s.textFont }
func (s *Setting) TextBoldFont() fyne.Resource       { return s.textBoldFont }
func (s *Setting) TextItalicFont() fyne.Resource     { return s.textItalicFont }
func (s *Setting) TextBoldItalicFont() fyne.Resource { return s.textBoldItalicFont }
func (s *Setting) TextMonospaceFont() fyne.Resource  { return s.textMonospaceFont }

func (s *Setting) SetPackageName(name string)               { s.packageName = name }
func (s *Setting) SetThemeStructName(name string)           { s.themeStructName = name }
func (s *Setting) SetBackgroundColor(c color.Color)         { s.backgroundColor = c }
func (s *Setting) SetButtonColor(c color.Color)             { s.buttonColor = c }
func (s *Setting) SetDisabledButtonColor(c color.Color)     { s.disabledButtonColor = c }
func (s *Setting) SetDisabledColor(c color.Color)           { s.disabledColor = c }
func (s *Setting) SetErrorColor(c color.Color)              { s.errorColor = c }
func (s *Setting) SetFocusColor(c color.Color)              { s.focusColor = c }
func (s *Setting) SetForegroundColor(c color.Color)         { s.foregroundColor = c }
func (s *Setting) SetHoverColor(c color.Color)              { s.hoverColor = c }
func (s *Setting) SetInputBackgroundColor(c color.Color)    { s.inputBackgroundColor = c }
func (s *Setting) SetPlaceHolderColor(c color.Color)        { s.placeHolderColor = c }
func (s *Setting) SetPressedColor(c color.Color)            { s.pressedColor = c }
func (s *Setting) SetPrimaryColor(c color.Color)            { s.primaryColor = c }
func (s *Setting) SetScrollBarColor(c color.Color)          { s.scrollBarColor = c }
func (s *Setting) SetShadowColor(c color.Color)             { s.shadowColor = c }
func (s *Setting) SetCaptionTextSize(size float32)          { s.captionTextSize = size }
func (s *Setting) SetInlineIconSize(size float32)           { s.inlineIconSize = size }
func (s *Setting) SetPaddingSize(size float32)              { s.paddingSize = size }
func (s *Setting) SetScrollBarSize(size float32)            { s.scrollBarSize = size }
func (s *Setting) SetScrollBarSmallSize(size float32)       { s.scrollBarSmallSize = size }
func (s *Setting) SetSeparatorThicknessSize(size float32)   { s.separatorThicknessSize = size }
func (s *Setting) SetTextSize(size float32)                 { s.textSize = size }
func (s *Setting) SetInputBorderSize(size float32)          { s.inputBorderSize = size }
func (s *Setting) SetTextFont(font fyne.Resource)           { s.textFont = font }
func (s *Setting) SetTextBoldFont(font fyne.Resource)       { s.textBoldFont = font }
func (s *Setting) SetTextItalicFont(font fyne.Resource)     { s.textItalicFont = font }
func (s *Setting) SetTextBoldItalicFont(font fyne.Resource) { s.textBoldItalicFont = font }
func (s *Setting) SetTextMonospaceFont(font fyne.Resource)  { s.textMonospaceFont = font }

func (s *Setting) ExportFontFile() bool     { return s.exportFontFile }
func (s *Setting) SetExportFontFile(b bool) { s.exportFontFile = b }
func (s *Setting) ExportForV2() bool        { return s.exportForV2 }
func (s *Setting) SetExportForV2(b bool)    { s.exportForV2 = b }

func (s *Setting) UpdateTheme(t fyne.Theme, v fyne.ThemeVariant) {
	s.SetBackgroundColor(t.Color(theme.ColorNameBackground, v))
	s.SetButtonColor(t.Color(theme.ColorNameButton, v))
	s.SetDisabledButtonColor(t.Color(theme.ColorNameDisabledButton, v))
	s.SetDisabledColor(t.Color(theme.ColorNameDisabled, v))
	s.SetErrorColor(t.Color(theme.ColorNameError, v))
	s.SetFocusColor(t.Color(theme.ColorNameFocus, v))
	s.SetForegroundColor(t.Color(theme.ColorNameForeground, v))
	s.SetHoverColor(t.Color(theme.ColorNameHover, v))
	s.SetInputBackgroundColor(t.Color(theme.ColorNameInputBackground, v))
	s.SetPlaceHolderColor(t.Color(theme.ColorNamePlaceHolder, v))
	s.SetPressedColor(t.Color(theme.ColorNamePressed, v))
	s.SetPrimaryColor(t.Color(theme.ColorNamePrimary, v))
	s.SetScrollBarColor(t.Color(theme.ColorNameScrollBar, v))
	s.SetShadowColor(t.Color(theme.ColorNameShadow, v))
	s.SetCaptionTextSize(t.Size(theme.SizeNameCaptionText))
	s.SetInlineIconSize(t.Size(theme.SizeNameInlineIcon))
	s.SetPaddingSize(t.Size(theme.SizeNamePadding))
	s.SetScrollBarSize(t.Size(theme.SizeNameScrollBar))
	s.SetScrollBarSmallSize(t.Size(theme.SizeNameScrollBarSmall))
	s.SetSeparatorThicknessSize(t.Size(theme.SizeNameSeparatorThickness))
	s.SetTextSize(t.Size(theme.SizeNameText))
	s.SetInputBorderSize(t.Size(theme.SizeNameInputBorder))
	s.SetTextFont(t.Font(fyne.TextStyle{}))
	s.SetTextBoldFont(t.Font(fyne.TextStyle{Bold: true}))
	s.SetTextItalicFont(t.Font(fyne.TextStyle{Italic: true}))
	s.SetTextBoldItalicFont(t.Font(fyne.TextStyle{Bold: true, Italic: true}))
	s.SetTextMonospaceFont(t.Font(fyne.TextStyle{Monospace: true}))
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
	case theme.ColorNameButton:
		return s.buttonColor
	case theme.ColorNameDisabledButton:
		return s.disabledButtonColor
	case theme.ColorNameDisabled:
		return s.disabledColor
	case theme.ColorNameError:
		return s.errorColor
	case theme.ColorNameFocus:
		return s.focusColor
	case theme.ColorNameForeground:
		return s.foregroundColor
	case theme.ColorNameHover:
		return s.hoverColor
	case theme.ColorNameInputBackground:
		return s.inputBackgroundColor
	case theme.ColorNamePlaceHolder:
		return s.placeHolderColor
	case theme.ColorNamePressed:
		return s.pressedColor
	case theme.ColorNamePrimary:
		return s.primaryColor
	case theme.ColorNameScrollBar:
		return s.scrollBarColor
	case theme.ColorNameShadow:
		return s.shadowColor
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
	case theme.SizeNameCaptionText:
		return s.captionTextSize
	case theme.SizeNameInlineIcon:
		return s.inlineIconSize
	case theme.SizeNamePadding:
		return s.paddingSize
	case theme.SizeNameScrollBar:
		return s.scrollBarSize
	case theme.SizeNameScrollBarSmall:
		return s.scrollBarSmallSize
	case theme.SizeNameSeparatorThickness:
		return s.separatorThicknessSize
	case theme.SizeNameText:
		return s.textSize
	case theme.SizeNameInputBorder:
		return s.inputBorderSize
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
