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
}

func NewSetting() *Setting {
	return &Setting{
		packageName:         defaultPackageName,
		themeStructName:     defaultThemeStructName,
		backgroundColor:     theme.DarkTheme().BackgroundColor(),
		buttonColor:         theme.DarkTheme().ButtonColor(),
		disabledButtonColor: theme.DarkTheme().DisabledButtonColor(),
		textColor:           theme.DarkTheme().TextColor(),
		disabledTextColor:   theme.DarkTheme().DisabledTextColor(),
		iconColor:           theme.DarkTheme().IconColor(),
		disabledIconColor:   theme.DarkTheme().DisabledIconColor(),
		hyperlinkColor:      theme.DarkTheme().HyperlinkColor(),
		placeHolderColor:    theme.DarkTheme().PlaceHolderColor(),
		primaryColor:        theme.DarkTheme().PrimaryColor(),
		hoverColor:          theme.DarkTheme().HoverColor(),
		focusColor:          theme.DarkTheme().FocusColor(),
		scrollBarColor:      theme.DarkTheme().ScrollBarColor(),
		shadowColor:         theme.DarkTheme().ShadowColor(),
		textSize:            theme.DarkTheme().TextSize(),
		textFont:            theme.DarkTheme().TextFont(),
		textBoldFont:        theme.DarkTheme().TextBoldFont(),
		textItalicFont:      theme.DarkTheme().TextItalicFont(),
		textBoldItalicFont:  theme.DarkTheme().TextBoldItalicFont(),
		textMonospaceFont:   theme.DarkTheme().TextMonospaceFont(),
		padding:             theme.DarkTheme().Padding(),
		iconInlineSize:      theme.DarkTheme().IconInlineSize(),
		scrollBarSize:       theme.DarkTheme().ScrollBarSize(),
		scrollBarSmallSize:  theme.DarkTheme().ScrollBarSmallSize(),
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
