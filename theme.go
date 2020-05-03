package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type themeSetting struct {
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

func newThemeSetting() *themeSetting {
	return &themeSetting{
		packageName:         "main",
		themeStructName:     "myTheme",
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

func (s *themeSetting) BackgroundColor() color.Color      { return s.backgroundColor }
func (s *themeSetting) ButtonColor() color.Color          { return s.buttonColor }
func (s *themeSetting) DisabledButtonColor() color.Color  { return s.disabledButtonColor }
func (s *themeSetting) TextColor() color.Color            { return s.textColor }
func (s *themeSetting) DisabledTextColor() color.Color    { return s.disabledTextColor }
func (s *themeSetting) IconColor() color.Color            { return s.iconColor }
func (s *themeSetting) DisabledIconColor() color.Color    { return s.disabledIconColor }
func (s *themeSetting) HyperlinkColor() color.Color       { return s.hyperlinkColor }
func (s *themeSetting) PlaceHolderColor() color.Color     { return s.placeHolderColor }
func (s *themeSetting) PrimaryColor() color.Color         { return s.primaryColor }
func (s *themeSetting) HoverColor() color.Color           { return s.hoverColor }
func (s *themeSetting) FocusColor() color.Color           { return s.focusColor }
func (s *themeSetting) ScrollBarColor() color.Color       { return s.scrollBarColor }
func (s *themeSetting) ShadowColor() color.Color          { return s.shadowColor }
func (s *themeSetting) TextSize() int                     { return s.textSize }
func (s *themeSetting) TextFont() fyne.Resource           { return s.textFont }
func (s *themeSetting) TextBoldFont() fyne.Resource       { return s.textBoldFont }
func (s *themeSetting) TextItalicFont() fyne.Resource     { return s.textItalicFont }
func (s *themeSetting) TextBoldItalicFont() fyne.Resource { return s.textBoldItalicFont }
func (s *themeSetting) TextMonospaceFont() fyne.Resource  { return s.textMonospaceFont }
func (s *themeSetting) Padding() int                      { return s.padding }
func (s *themeSetting) IconInlineSize() int               { return s.iconInlineSize }
func (s *themeSetting) ScrollBarSize() int                { return s.scrollBarSize }
func (s *themeSetting) ScrollBarSmallSize() int           { return s.scrollBarSmallSize }

func (s *themeSetting) SetBackgroundColor(c color.Color)         { s.backgroundColor = c }
func (s *themeSetting) SetButtonColor(c color.Color)             { s.buttonColor = c }
func (s *themeSetting) SetDisabledButtonColor(c color.Color)     { s.disabledButtonColor = c }
func (s *themeSetting) SetTextColor(c color.Color)               { s.textColor = c }
func (s *themeSetting) SetDisabledTextColor(c color.Color)       { s.disabledTextColor = c }
func (s *themeSetting) SetIconColor(c color.Color)               { s.iconColor = c }
func (s *themeSetting) SetDisabledIconColor(c color.Color)       { s.disabledIconColor = c }
func (s *themeSetting) SetHyperlinkColor(c color.Color)          { s.hyperlinkColor = c }
func (s *themeSetting) SetPlaceHolderColor(c color.Color)        { s.placeHolderColor = c }
func (s *themeSetting) SetPrimaryColor(c color.Color)            { s.primaryColor = c }
func (s *themeSetting) SetHoverColor(c color.Color)              { s.hoverColor = c }
func (s *themeSetting) SetFocusColor(c color.Color)              { s.focusColor = c }
func (s *themeSetting) SetScrollBarColor(c color.Color)          { s.scrollBarColor = c }
func (s *themeSetting) SetShadowColor(c color.Color)             { s.shadowColor = c }
func (s *themeSetting) SetTextSize(size int)                     { s.textSize = size }
func (s *themeSetting) SetTextFont(font fyne.Resource)           { s.textFont = font }
func (s *themeSetting) SetTextBoldFont(font fyne.Resource)       { s.textBoldFont = font }
func (s *themeSetting) SetTextItalicFont(font fyne.Resource)     { s.textItalicFont = font }
func (s *themeSetting) SetTextBoldItalicFont(font fyne.Resource) { s.textBoldItalicFont = font }
func (s *themeSetting) SetTextMonospaceFont(font fyne.Resource)  { s.textMonospaceFont = font }
func (s *themeSetting) SetPadding(pad int)                       { s.padding = pad }
func (s *themeSetting) SetIconInlineSize(size int)               { s.iconInlineSize = size }
func (s *themeSetting) SetScrollBarSize(size int)                { s.scrollBarSize = size }
func (s *themeSetting) SetScrollBarSmallSize(size int)           { s.scrollBarSmallSize = size }
