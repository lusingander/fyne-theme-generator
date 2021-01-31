package ui

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	ft "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/lusingander/colorpicker"
	"github.com/lusingander/fyne-theme-generator/internal/theme"
)

type configPanel struct {
	panel   fyne.CanvasObject
	parent  fyne.Window
	current *theme.Setting
	refresh func()

	backgroundColorSelector      *colorSelector
	buttonColorSelector          *colorSelector
	disabledButtonColorSelector  *colorSelector
	disabledColorSelector        *colorSelector
	errorColorSelector           *colorSelector
	focusColorSelector           *colorSelector
	foregroundColorSelector      *colorSelector
	hoverColorSelector           *colorSelector
	inputBackgroundColorSelector *colorSelector
	placeHolderColorSelector     *colorSelector
	pressedColorSelector         *colorSelector
	primaryColorSelector         *colorSelector
	scrollBarColorSelector       *colorSelector
	shadowColorSelector          *colorSelector

	captionTextSizeSelector        *floatSelector
	inlineIconSizeSelector         *floatSelector
	paddingSizeSelector            *floatSelector
	scrollBarSizeSelector          *floatSelector
	scrollBarSmallSizeSelector     *floatSelector
	separatorThicknessSizeSelector *floatSelector
	textSizeSelector               *floatSelector
	inputBorderSizeSelector        *floatSelector

	textFontSelector           *fontFilepathSelector
	textBoldFontSelector       *fontFilepathSelector
	textItalicFontSelector     *fontFilepathSelector
	textBoldItalicFontSelector *fontFilepathSelector
	textMonospaceFontSelector  *fontFilepathSelector
}

func (u *ui) newConfigPanel() *configPanel {
	p := &configPanel{
		parent:  u.window,
		current: u.current,
		refresh: u.refresh,
	}
	p.build()
	return p
}

func (p *configPanel) applyCurrentTheme() {
	p.backgroundColorSelector.setColor(p.current.BackgroundColor())
	p.buttonColorSelector.setColor(p.current.ButtonColor())
	p.disabledButtonColorSelector.setColor(p.current.DisabledButtonColor())
	p.disabledColorSelector.setColor(p.current.DisabledColor())
	p.errorColorSelector.setColor(p.current.ErrorColor())
	p.focusColorSelector.setColor(p.current.FocusColor())
	p.foregroundColorSelector.setColor(p.current.ForegroundColor())
	p.hoverColorSelector.setColor(p.current.HoverColor())
	p.inputBackgroundColorSelector.setColor(p.current.InputBackgroundColor())
	p.placeHolderColorSelector.setColor(p.current.PlaceHolderColor())
	p.pressedColorSelector.setColor(p.current.PressedColor())
	p.primaryColorSelector.setColor(p.current.PrimaryColor())
	p.scrollBarColorSelector.setColor(p.current.ScrollBarColor())
	p.shadowColorSelector.setColor(p.current.ShadowColor())

	p.captionTextSizeSelector.setValue(p.current.CaptionTextSize())
	p.inlineIconSizeSelector.setValue(p.current.InlineIconSize())
	p.paddingSizeSelector.setValue(p.current.PaddingSize())
	p.scrollBarSizeSelector.setValue(p.current.ScrollBarSize())
	p.scrollBarSmallSizeSelector.setValue(p.current.ScrollBarSmallSize())
	p.separatorThicknessSizeSelector.setValue(p.current.SeparatorThicknessSize())
	p.textSizeSelector.setValue(p.current.TextSize())
	p.inputBorderSizeSelector.setValue(p.current.InputBorderSize())

	p.textFontSelector.setValue(p.current.TextFont().Name())
	p.textBoldFontSelector.setValue(p.current.TextBoldFont().Name())
	p.textItalicFontSelector.setValue(p.current.TextItalicFont().Name())
	p.textBoldItalicFontSelector.setValue(p.current.TextBoldItalicFont().Name())
	p.textMonospaceFontSelector.setValue(p.current.TextMonospaceFont().Name())
}

func (p *configPanel) build() {
	confs := p.configures(p.current)
	l := len(confs) / 2
	if l%2 != 0 {
		l++
	}
	p.panel = fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			confs[:l]...,
		),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			confs[l:]...,
		),
	)
}

func (p *configPanel) configures(ts *theme.Setting) []fyne.CanvasObject {
	p.backgroundColorSelector = p.newColorSelector(ts.BackgroundColor(), ts.SetBackgroundColor)
	p.buttonColorSelector = p.newColorSelector(ts.ButtonColor(), ts.SetButtonColor)
	p.disabledButtonColorSelector = p.newColorSelector(ts.DisabledButtonColor(), ts.SetDisabledButtonColor)
	p.disabledColorSelector = p.newColorSelector(ts.DisabledColor(), ts.SetDisabledColor)
	p.errorColorSelector = p.newColorSelector(ts.ErrorColor(), ts.SetErrorColor)
	p.focusColorSelector = p.newColorSelector(ts.FocusColor(), ts.SetFocusColor)
	p.foregroundColorSelector = p.newColorSelector(ts.ForegroundColor(), ts.SetForegroundColor)
	p.hoverColorSelector = p.newColorSelector(ts.HoverColor(), ts.SetHoverColor)
	p.inputBackgroundColorSelector = p.newColorSelector(ts.InputBackgroundColor(), ts.SetInputBackgroundColor)
	p.placeHolderColorSelector = p.newColorSelector(ts.PlaceHolderColor(), ts.SetPlaceHolderColor)
	p.pressedColorSelector = p.newColorSelector(ts.PressedColor(), ts.SetPressedColor)
	p.primaryColorSelector = p.newColorSelector(ts.PrimaryColor(), ts.SetPrimaryColor)
	p.scrollBarColorSelector = p.newColorSelector(ts.ScrollBarColor(), ts.SetScrollBarColor)
	p.shadowColorSelector = p.newColorSelector(ts.ShadowColor(), ts.SetShadowColor)

	p.captionTextSizeSelector = p.newFloatSelector(ts.CaptionTextSize(), ts.SetCaptionTextSize)
	p.inlineIconSizeSelector = p.newFloatSelector(ts.InlineIconSize(), ts.SetInlineIconSize)
	p.paddingSizeSelector = p.newFloatSelector(ts.PaddingSize(), ts.SetPaddingSize)
	p.scrollBarSizeSelector = p.newFloatSelector(ts.ScrollBarSize(), ts.SetScrollBarSize)
	p.scrollBarSmallSizeSelector = p.newFloatSelector(ts.ScrollBarSmallSize(), ts.SetScrollBarSmallSize)
	p.separatorThicknessSizeSelector = p.newFloatSelector(ts.SeparatorThicknessSize(), ts.SetSeparatorThicknessSize)
	p.textSizeSelector = p.newFloatSelector(ts.TextSize(), ts.SetTextSize)
	p.inputBorderSizeSelector = p.newFloatSelector(ts.InputBorderSize(), ts.SetInputBorderSize)

	p.textFontSelector = p.newFontFilepathSelector(ts.TextFont().Name(), ts.SetTextFont)
	p.textBoldFontSelector = p.newFontFilepathSelector(ts.TextBoldFont().Name(), ts.SetTextBoldFont)
	p.textItalicFontSelector = p.newFontFilepathSelector(ts.TextItalicFont().Name(), ts.SetTextItalicFont)
	p.textBoldItalicFontSelector = p.newFontFilepathSelector(ts.TextBoldItalicFont().Name(), ts.SetTextBoldItalicFont)
	p.textMonospaceFontSelector = p.newFontFilepathSelector(ts.TextMonospaceFont().Name(), ts.SetTextMonospaceFont)

	cs := make([]fyne.CanvasObject, 0)
	cs = append(cs, colorConfigure("Background color", p.backgroundColorSelector)...)
	cs = append(cs, colorConfigure("Button color", p.buttonColorSelector)...)
	cs = append(cs, colorConfigure("DisabledButton color", p.disabledButtonColorSelector)...)
	cs = append(cs, colorConfigure("Disabled color", p.disabledColorSelector)...)
	cs = append(cs, colorConfigure("Error color", p.errorColorSelector)...)
	cs = append(cs, colorConfigure("Focus color", p.focusColorSelector)...)
	cs = append(cs, colorConfigure("Foreground color", p.foregroundColorSelector)...)
	cs = append(cs, colorConfigure("Hover color", p.hoverColorSelector)...)
	cs = append(cs, colorConfigure("InputBackground color", p.inputBackgroundColorSelector)...)
	cs = append(cs, colorConfigure("PlaceHolder color", p.placeHolderColorSelector)...)
	cs = append(cs, colorConfigure("Pressed color", p.pressedColorSelector)...)
	cs = append(cs, colorConfigure("Primary color", p.primaryColorSelector)...)
	cs = append(cs, colorConfigure("ScrollBar color", p.scrollBarColorSelector)...)
	cs = append(cs, colorConfigure("Shadow color", p.shadowColorSelector)...)
	cs = append(cs, floatConfigure("CaptionText size", p.captionTextSizeSelector)...)
	cs = append(cs, floatConfigure("InlineIcon size", p.inlineIconSizeSelector)...)
	cs = append(cs, floatConfigure("Padding size", p.paddingSizeSelector)...)
	cs = append(cs, floatConfigure("ScrollBar size", p.scrollBarSizeSelector)...)
	cs = append(cs, floatConfigure("ScrollBarSmall size", p.scrollBarSmallSizeSelector)...)
	cs = append(cs, floatConfigure("SeparatorThickness size", p.separatorThicknessSizeSelector)...)
	cs = append(cs, floatConfigure("Text size", p.textSizeSelector)...)
	cs = append(cs, floatConfigure("InputBorder size", p.inputBorderSizeSelector)...)
	cs = append(cs, fontFilepathConfigure("Text font", p.textFontSelector)...)
	cs = append(cs, fontFilepathConfigure("TextBold font", p.textBoldFontSelector)...)
	cs = append(cs, fontFilepathConfigure("TextItalic font", p.textItalicFontSelector)...)
	cs = append(cs, fontFilepathConfigure("TextBoldItalic font", p.textBoldItalicFontSelector)...)
	cs = append(cs, fontFilepathConfigure("TextMonospace font", p.textMonospaceFontSelector)...)
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

func floatConfigure(label string, selector *floatSelector) []fyne.CanvasObject {
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

func fontFilepathConfigure(label string, selector *fontFilepathSelector) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		container.NewBorder(
			nil,
			nil,
			selector.button,
			nil,
			selector.entry,
		),
	}
}

type colorSelector struct {
	entry        *widget.Entry
	rect         colorpicker.PickerOpenWidget
	tmp          color.Color
	update       func(color.Color)
	sampleWidget fyne.CanvasObject
	refresh      func()
}

func (p *configPanel) newColorSelector(defaultColor color.Color, update func(color.Color)) *colorSelector {
	entry := &widget.Entry{TextStyle: fyne.TextStyle{Monospace: true}}
	rect := colorpicker.NewColorSelectModalRect(p.parent, fyne.NewSize(20, 20), defaultColor)
	rect.SetPickerStyle(colorpicker.StyleHueCircle)
	selector := &colorSelector{
		entry:   entry,
		rect:    rect,
		tmp:     defaultColor,
		update:  update,
		refresh: p.refresh,
	}
	selector.setColorAndRefresh(defaultColor)
	rect.SetOnChange(selector.setColorAndRefresh)
	entry.OnChanged = func(s string) {
		var r, g, b, a uint8
		l := len(s)
		if l > 9 {
			selector.setColorAndRefresh(selector.tmp)
		} else if _, err := fmt.Sscanf(s, "#%02x%02x%02x%02x", &r, &g, &b, &a); l == 9 && err == nil {
			selector.setColorAndRefresh(color.NRGBA{r, g, b, a})
		}
	}
	return selector
}

func (c *colorSelector) setColorAndRefresh(clr color.Color) {
	c.setColor(clr)
	c.update(clr)
	c.refresh()
}

func (c *colorSelector) setColor(clr color.Color) {
	c.tmp = clr
	c.entry.SetText(hexColorString(clr))
	c.rect.SetColor(clr)
}

func hexColorString(c color.Color) string {
	rgba, _ := c.(color.NRGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B, rgba.A)
}

type floatSelector struct {
	entry *widget.Entry
}

func (p *configPanel) newFloatSelector(defaultValue float32, update func(float32)) *floatSelector {
	entry := &widget.Entry{TextStyle: fyne.TextStyle{Monospace: true}}
	selector := &floatSelector{
		entry: entry,
	}
	selector.setValue(defaultValue)
	entry.OnChanged = func(s string) {
		if f, err := strconv.ParseFloat(s, 32); err == nil {
			update(float32(f))
			p.refresh()
		}
	}
	return selector
}

func (s *floatSelector) setValue(v float32) {
	s.entry.SetText(strconv.FormatFloat(float64(v), 'f', 2, 32))
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

func (s *readonlyStringSelector) setValue(v string) {
	s.str.SetText(v)
}

type fontFilepathSelector struct {
	parent   fyne.Window
	entry    *widget.Entry
	button   *widget.Button
	filepath string
	update   func(fyne.Resource)
	refresh  func()
}

func (p *configPanel) newFontFilepathSelector(defaultValue string, update func(fyne.Resource)) *fontFilepathSelector {
	selector := &fontFilepathSelector{
		parent:  p.parent,
		update:  update,
		refresh: p.refresh,
	}
	selector.entry = widget.NewEntry()
	selector.entry.Disable()
	selector.entry.SetText(defaultValue)
	selector.button = widget.NewButtonWithIcon("", ft.FolderOpenIcon(), selector.openFileDialog)
	return selector
}

func (s *fontFilepathSelector) setValue(v string) {
	s.entry.SetText(v)
}

func (s *fontFilepathSelector) openFileDialog() {
	d := dialog.NewFileOpen(s.loadFontfileWrapper, s.parent)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".ttf", ".otf"}))
	d.Show()
}

func (s *fontFilepathSelector) loadFontfileWrapper(reader fyne.URIReadCloser, err error) {
	if err == nil && reader != nil {
		err = s.loadFontfile(reader)
	}
	if err != nil {
		dialog.ShowError(err, s.parent)
	}
}

func (s *fontFilepathSelector) loadFontfile(reader fyne.URIReadCloser) error {
	defer reader.Close()
	path := reader.URI().String()[7:] // `file://`
	res, err := fyne.LoadResourceFromPath(path)
	if err != nil {
		return err
	}
	s.setValue(res.Name())
	s.update(res)
	s.refresh()
	return nil
}
