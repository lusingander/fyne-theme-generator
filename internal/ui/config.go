package ui

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/lusingander/colorpicker"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

type configPanel struct {
	panel   fyne.CanvasObject
	parent  fyne.Window
	current *theme.Setting
	reflesh func()

	backgroundColorSelector     *colorSelector
	buttonColorSelector         *colorSelector
	disabledButtonColorSelector *colorSelector
	textColorSelector           *colorSelector
	disabledTextColorSelector   *colorSelector
	iconColorSelector           *colorSelector
	disabledIconColorSelector   *colorSelector
	hyperlinkColorSelector      *colorSelector
	placeHolderColorSelector    *colorSelector
	primaryColorSelector        *colorSelector
	hoverColorSelector          *colorSelector
	focusColorSelector          *colorSelector
	scrollBarColorSelector      *colorSelector
	shadowColorSelector         *colorSelector

	textSizeSelector           *intSelector
	textFontSelector           *readonlyStringSelector
	textBoldFontSelector       *readonlyStringSelector
	textItalicFontSelector     *readonlyStringSelector
	textBoldItalicFontSelector *readonlyStringSelector
	textMonospaceFontSelector  *readonlyStringSelector

	paddingSelector            *intSelector
	iconInlineSizeSelector     *intSelector
	scrollBarSizeSelector      *intSelector
	scrollBarSmallSizeSelector *intSelector
}

func (u *ui) newConfigPanel() *configPanel {
	p := &configPanel{
		parent:  u.window,
		current: u.current,
		reflesh: u.Reflesh,
	}
	p.build()
	return p
}

func (p *configPanel) build() {
	confs := p.configures(p.current)
	p.panel = fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			confs[:len(confs)/2]...,
		),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			confs[len(confs)/2:]...,
		),
	)
}

func (p *configPanel) configures(ts *theme.Setting) []fyne.CanvasObject {
	p.backgroundColorSelector = p.newColorSelector(ts.BackgroundColor(), ts.SetBackgroundColor)
	p.buttonColorSelector = p.newColorSelector(ts.ButtonColor(), ts.SetButtonColor)
	p.disabledButtonColorSelector = p.newColorSelector(ts.DisabledButtonColor(), ts.SetDisabledButtonColor)
	p.textColorSelector = p.newColorSelector(ts.TextColor(), ts.SetTextColor)
	p.disabledTextColorSelector = p.newColorSelector(ts.DisabledTextColor(), ts.SetDisabledTextColor)
	p.iconColorSelector = p.newColorSelector(ts.IconColor(), ts.SetIconColor)
	p.disabledIconColorSelector = p.newColorSelector(ts.DisabledIconColor(), ts.SetDisabledIconColor)
	p.hyperlinkColorSelector = p.newColorSelector(ts.HyperlinkColor(), ts.SetHyperlinkColor)
	p.placeHolderColorSelector = p.newColorSelector(ts.PlaceHolderColor(), ts.SetPlaceHolderColor)
	p.primaryColorSelector = p.newColorSelector(ts.PrimaryColor(), ts.SetPrimaryColor)
	p.hoverColorSelector = p.newColorSelector(ts.HoverColor(), ts.SetHoverColor)
	p.focusColorSelector = p.newColorSelector(ts.FocusColor(), ts.SetFocusColor)
	p.scrollBarColorSelector = p.newColorSelector(ts.ScrollBarColor(), ts.SetScrollBarColor)
	p.shadowColorSelector = p.newColorSelector(ts.ShadowColor(), ts.SetShadowColor)
	p.textSizeSelector = p.newIntSelector(ts.TextSize(), ts.SetTextSize)
	p.textFontSelector = p.newReadonlyStringSelector(ts.TextFont().Name())
	p.textBoldFontSelector = p.newReadonlyStringSelector(ts.TextBoldFont().Name())
	p.textItalicFontSelector = p.newReadonlyStringSelector(ts.TextItalicFont().Name())
	p.textBoldItalicFontSelector = p.newReadonlyStringSelector(ts.TextBoldItalicFont().Name())
	p.textMonospaceFontSelector = p.newReadonlyStringSelector(ts.TextMonospaceFont().Name())
	p.paddingSelector = p.newIntSelector(ts.Padding(), ts.SetPadding)
	p.iconInlineSizeSelector = p.newIntSelector(ts.IconInlineSize(), ts.SetIconInlineSize)
	p.scrollBarSizeSelector = p.newIntSelector(ts.ScrollBarSize(), ts.SetScrollBarSize)
	p.scrollBarSmallSizeSelector = p.newIntSelector(ts.ScrollBarSmallSize(), ts.SetScrollBarSmallSize)

	cs := make([]fyne.CanvasObject, 0)
	cs = append(cs, colorConfigure("Background color", p.backgroundColorSelector)...)
	cs = append(cs, colorConfigure("Button color", p.buttonColorSelector)...)
	cs = append(cs, colorConfigure("Disable button color", p.disabledButtonColorSelector)...)
	cs = append(cs, colorConfigure("Text color", p.textColorSelector)...)
	cs = append(cs, colorConfigure("Disable text color", p.disabledTextColorSelector)...)
	cs = append(cs, colorConfigure("Icon color", p.iconColorSelector)...)
	cs = append(cs, colorConfigure("Disable icon color", p.disabledIconColorSelector)...)
	cs = append(cs, colorConfigure("Hyperlink color", p.hyperlinkColorSelector)...)
	cs = append(cs, colorConfigure("Placeholder color", p.placeHolderColorSelector)...)
	cs = append(cs, colorConfigure("Primary color", p.primaryColorSelector)...)
	cs = append(cs, colorConfigure("Hover color", p.hoverColorSelector)...)
	cs = append(cs, colorConfigure("Focus color", p.focusColorSelector)...)
	cs = append(cs, colorConfigure("Scroll bar color", p.scrollBarColorSelector)...)
	cs = append(cs, colorConfigure("Shadow color", p.shadowColorSelector)...)
	cs = append(cs, intConfigure("Text size", p.textSizeSelector)...)
	cs = append(cs, readonlyStringConfigure("Text font", p.textFontSelector)...)
	cs = append(cs, readonlyStringConfigure("Text bold font", p.textBoldFontSelector)...)
	cs = append(cs, readonlyStringConfigure("Text italic font", p.textItalicFontSelector)...)
	cs = append(cs, readonlyStringConfigure("Text bold italic font", p.textBoldItalicFontSelector)...)
	cs = append(cs, readonlyStringConfigure("Text monospace font", p.textMonospaceFontSelector)...)
	cs = append(cs, intConfigure("Padding", p.paddingSelector)...)
	cs = append(cs, intConfigure("Icon inline size", p.iconInlineSizeSelector)...)
	cs = append(cs, intConfigure("Scroll bar size", p.scrollBarSizeSelector)...)
	cs = append(cs, intConfigure("Scroll bar small size", p.scrollBarSmallSizeSelector)...)
	return cs
}

func colorConfigure(label string, selector *colorSelector) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			selector.rect,
			selector.entry,
		),
	}
}

func intConfigure(label string, selector *intSelector) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		selector.entry,
	}
}

func readonlyStringConfigure(label string, selector *readonlyStringSelector) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		selector.str,
	}
}

type colorSelector struct {
	entry        *widget.Entry
	rect         colorpicker.PickerOpenWidget
	tmp          color.Color
	update       func(color.Color)
	sampleWidget fyne.CanvasObject
	reflesh      func()
}

func (p *configPanel) newColorSelector(defaultColor color.Color, update func(color.Color)) *colorSelector {
	entry := &widget.Entry{}
	rect := colorpicker.NewColorSelectModalRect(p.parent, fyne.NewSize(20, 20), defaultColor)
	rect.SetPickerStyle(colorpicker.StyleCircle)
	selector := &colorSelector{
		entry:   entry,
		rect:    rect,
		tmp:     defaultColor,
		update:  update,
		reflesh: p.reflesh,
	}
	selector.setColor(defaultColor)
	// colorpicker doesn't currently consider alpha...
	rect.SetOnChange(selector.setColorKeepAlpha)
	entry.OnChanged = func(s string) {
		var r, g, b, a uint8
		l := len(s)
		if l > 9 {
			selector.setColor(selector.tmp)
		} else if _, err := fmt.Sscanf(s, "#%02x%02x%02x%02x", &r, &g, &b, &a); l == 9 && err == nil {
			selector.setColor(color.RGBA{r, g, b, a})
		}
	}
	return selector
}

func (c *colorSelector) setColorKeepAlpha(clr color.Color) {
	r, g, b, _ := clr.RGBA()
	_, _, _, a := c.tmp.RGBA()
	c.setColor(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
}

func (c *colorSelector) setColor(clr color.Color) {
	c.tmp = clr
	c.entry.SetText(hexColorString(clr))
	c.rect.SetColor(clr)
	c.update(clr)
	c.reflesh()
}

func hexColorString(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B, rgba.A)
}

type intSelector struct {
	entry *widget.Entry
}

func (p *configPanel) newIntSelector(defaultValue int, update func(int)) *intSelector {
	entry := &widget.Entry{}
	entry.SetText(strconv.Itoa(defaultValue))
	entry.OnChanged = func(s string) {
		if n, err := strconv.Atoi(s); err == nil {
			update(n)
			p.reflesh()
		}
	}
	return &intSelector{
		entry: entry,
	}
}

type readonlyStringSelector struct {
	str *widget.Label
}

func (p *configPanel) newReadonlyStringSelector(defaultValue string) *readonlyStringSelector {
	str := widget.NewLabel(defaultValue)
	return &readonlyStringSelector{
		str: str,
	}
}
