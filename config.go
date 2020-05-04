package main

import (
	"fmt"
	"image/color"
	"net/url"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
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

func newColorSelector(defaultColor color.Color, update func(color.Color), sampleWidget fyne.CanvasObject) *colorSelector {
	entry := &widget.Entry{}
	rect := colorpicker.NewColorSelectModalRect(mainWindow, fyne.NewSize(20, 20), defaultColor)
	selector := &colorSelector{
		entry:        entry,
		rect:         rect,
		tmp:          defaultColor,
		update:       update,
		sampleWidget: sampleWidget,
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
		} else if _, err := fmt.Sscanf(s, "#%1x%1x%1x%1x", &r, &g, &b, &a); l == 5 && err == nil {
			selector.setColor(color.RGBA{r * 17, g * 17, b * 17, a * 17})
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

func colorConfigure(label string, sampleWidget fyne.CanvasObject, defaultColor color.Color, update func(color.Color)) []fyne.CanvasObject {
	cs := newColorSelector(defaultColor, update, sampleWidget)
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			cs.rect,
			cs.entry,
		),
		sampleWidget,
	}
}

func intConfigure(label string, sampleWidget fyne.CanvasObject, defaultValue int, update func(int)) []fyne.CanvasObject {
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
		sampleWidget,
	}
}

func readonlyStringConfigure(label string, sampleWidget fyne.CanvasObject, defaultValue string) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabel(label),
		widget.NewLabel(defaultValue),
		sampleWidget,
	}
}

func configures(ts *themeSetting) []fyne.CanvasObject {
	cs := make([]fyne.CanvasObject, 0)
	cs = append(cs, colorConfigure("Background color", empty(), ts.BackgroundColor(), ts.SetBackgroundColor)...)
	cs = append(cs, colorConfigure("Button color", dummyButton(true), ts.ButtonColor(), ts.SetButtonColor)...)
	cs = append(cs, colorConfigure("Disable button color", dummyButton(false), ts.DisabledButtonColor(), ts.SetDisabledButtonColor)...)
	cs = append(cs, colorConfigure("Text color", dummyText(), ts.TextColor(), ts.SetTextColor)...)
	cs = append(cs, colorConfigure("Disable text color", empty(), ts.DisabledTextColor(), ts.SetDisabledTextColor)...)
	cs = append(cs, colorConfigure("Icon color", dummyIcon(), ts.IconColor(), ts.SetIconColor)...)
	cs = append(cs, colorConfigure("Disable icon color", empty(), ts.DisabledIconColor(), ts.SetDisabledIconColor)...)
	cs = append(cs, colorConfigure("Hyperlink color", dummyHyperlink(), ts.HyperlinkColor(), ts.SetHyperlinkColor)...)
	cs = append(cs, colorConfigure("Placeholder color", dummyEntry(), ts.PlaceHolderColor(), ts.SetPlaceHolderColor)...)
	cs = append(cs, colorConfigure("Primary color", empty(), ts.PrimaryColor(), ts.SetPrimaryColor)...)
	cs = append(cs, colorConfigure("Hover color", empty(), ts.HoverColor(), ts.SetHoverColor)...)
	cs = append(cs, colorConfigure("Focus color", empty(), ts.FocusColor(), ts.SetFocusColor)...)
	cs = append(cs, colorConfigure("Scroll bar color", empty(), ts.ScrollBarColor(), ts.SetScrollBarColor)...)
	cs = append(cs, colorConfigure("Shadow color", empty(), ts.ShadowColor(), ts.SetShadowColor)...)
	cs = append(cs, intConfigure("Text size", empty(), ts.TextSize(), ts.SetTextSize)...)
	cs = append(cs, readonlyStringConfigure("Text font", dummyText(), ts.TextFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold font", dummyBoldText(), ts.TextBoldFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text italic font", dummyItalicText(), ts.TextItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold italic font", dummyBoldItalicText(), ts.TextBoldItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text monospace font", dummyMonospaceText(), ts.TextMonospaceFont().Name())...)
	cs = append(cs, intConfigure("Padding", empty(), ts.Padding(), ts.SetPadding)...)
	cs = append(cs, intConfigure("Icon inline size", empty(), ts.IconInlineSize(), ts.SetIconInlineSize)...)
	cs = append(cs, intConfigure("Scroll bar size", empty(), ts.ScrollBarSize(), ts.SetScrollBarSize)...)
	cs = append(cs, intConfigure("Scroll bar small size", empty(), ts.ScrollBarSmallSize(), ts.SetScrollBarSmallSize)...)
	return cs
}

func empty() fyne.CanvasObject {
	return layout.NewSpacer()
}

func dummyButton(enable bool) fyne.CanvasObject {
	button := widget.NewButton("Button sample", func() {})
	if !enable {
		button.Disable()
	}
	return button
}

func dummyText() fyne.CanvasObject {
	return dummyStyledText(false, false, false)
}

func dummyBoldText() fyne.CanvasObject {
	return dummyStyledText(true, false, false)
}

func dummyItalicText() fyne.CanvasObject {
	return dummyStyledText(false, true, false)
}

func dummyBoldItalicText() fyne.CanvasObject {
	return dummyStyledText(true, true, false)
}

func dummyMonospaceText() fyne.CanvasObject {
	return dummyStyledText(false, false, true)
}

func dummyStyledText(bold, italic, monospace bool) fyne.CanvasObject {
	text := widget.NewLabel("Text sample")
	text.TextStyle = fyne.TextStyle{Bold: bold, Italic: italic, Monospace: monospace}
	return text
}

func dummyIcon() fyne.CanvasObject {
	icon := widget.NewIcon(theme.DocumentSaveIcon())
	return icon
}

func dummyHyperlink() fyne.CanvasObject {
	url, _ := url.Parse("https://fyne.io/")
	hyperlink := widget.NewHyperlink("Hyperlink sample", url)
	return hyperlink
}

func dummyEntry() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry sample")
	return entry
}
