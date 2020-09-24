package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/lusingander/colorpicker"
)

type colorSelector struct {
	entry        *widget.Entry
	rect         colorpicker.PickerOpenWidget
	tmp          color.Color
	update       func(color.Color)
	sampleWidget fyne.CanvasObject
}

func newColorSelector(defaultColor color.Color, update func(color.Color)) *colorSelector {
	entry := &widget.Entry{}
	rect := colorpicker.NewColorSelectModalRect(mainWindow, fyne.NewSize(20, 20), defaultColor)
	selector := &colorSelector{
		entry:  entry,
		rect:   rect,
		tmp:    defaultColor,
		update: update,
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
	reflesh()
}

func hexColorString(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B, rgba.A)
}

func colorConfigure(label string, defaultColor color.Color, update func(color.Color)) []fyne.CanvasObject {
	cs := newColorSelector(defaultColor, update)
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			cs.rect,
			cs.entry,
		),
	}
}

func intConfigure(label string, defaultValue int, update func(int)) []fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetText(strconv.Itoa(defaultValue))
	entry.OnChanged = func(s string) {
		if n, err := strconv.Atoi(s); err == nil {
			update(n)
			reflesh()
		}
	}
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		entry,
	}
}

func readonlyStringConfigure(label string, defaultValue string) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		widget.NewLabel(defaultValue),
	}
}

func configures(ts *themeSetting) []fyne.CanvasObject {
	cs := make([]fyne.CanvasObject, 0)
	cs = append(cs, colorConfigure("Background color", ts.BackgroundColor(), ts.SetBackgroundColor)...)
	cs = append(cs, colorConfigure("Button color", ts.ButtonColor(), ts.SetButtonColor)...)
	cs = append(cs, colorConfigure("Disable button color", ts.DisabledButtonColor(), ts.SetDisabledButtonColor)...)
	cs = append(cs, colorConfigure("Text color", ts.TextColor(), ts.SetTextColor)...)
	cs = append(cs, colorConfigure("Disable text color", ts.DisabledTextColor(), ts.SetDisabledTextColor)...)
	cs = append(cs, colorConfigure("Icon color", ts.IconColor(), ts.SetIconColor)...)
	cs = append(cs, colorConfigure("Disable icon color", ts.DisabledIconColor(), ts.SetDisabledIconColor)...)
	cs = append(cs, colorConfigure("Hyperlink color", ts.HyperlinkColor(), ts.SetHyperlinkColor)...)
	cs = append(cs, colorConfigure("Placeholder color", ts.PlaceHolderColor(), ts.SetPlaceHolderColor)...)
	cs = append(cs, colorConfigure("Primary color", ts.PrimaryColor(), ts.SetPrimaryColor)...)
	cs = append(cs, colorConfigure("Hover color", ts.HoverColor(), ts.SetHoverColor)...)
	cs = append(cs, colorConfigure("Focus color", ts.FocusColor(), ts.SetFocusColor)...)
	cs = append(cs, colorConfigure("Scroll bar color", ts.ScrollBarColor(), ts.SetScrollBarColor)...)
	cs = append(cs, colorConfigure("Shadow color", ts.ShadowColor(), ts.SetShadowColor)...)
	cs = append(cs, intConfigure("Text size", ts.TextSize(), ts.SetTextSize)...)
	cs = append(cs, readonlyStringConfigure("Text font", ts.TextFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold font", ts.TextBoldFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text italic font", ts.TextItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold italic font", ts.TextBoldItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text monospace font", ts.TextMonospaceFont().Name())...)
	cs = append(cs, intConfigure("Padding", ts.Padding(), ts.SetPadding)...)
	cs = append(cs, intConfigure("Icon inline size", ts.IconInlineSize(), ts.SetIconInlineSize)...)
	cs = append(cs, intConfigure("Scroll bar size", ts.ScrollBarSize(), ts.SetScrollBarSize)...)
	cs = append(cs, intConfigure("Scroll bar small size", ts.ScrollBarSmallSize(), ts.SetScrollBarSmallSize)...)
	return cs
}

func config() fyne.CanvasObject {
	confs := configures(currentThemeSetting)
	return fyne.NewContainerWithLayout(
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
