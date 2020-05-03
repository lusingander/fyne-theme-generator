package main

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"os"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type colorSelector struct {
	entry        *widget.Entry
	rect         *canvas.Rectangle
	tmp          color.Color
	update       func(color.Color)
	sampleWidget fyne.CanvasObject
}

func newColorSelector(defaultColor color.Color, update func(color.Color), sampleWidget fyne.CanvasObject) *colorSelector {
	entry := &widget.Entry{PlaceHolder: "#000000"}
	rect := &canvas.Rectangle{FillColor: color.RGBA{0, 0, 0, 255}}
	rect.SetMinSize(fyne.NewSize(20, 20))
	selector := &colorSelector{
		entry:        entry,
		rect:         rect,
		tmp:          defaultColor,
		update:       update,
		sampleWidget: sampleWidget,
	}
	selector.setColor(defaultColor)
	entry.OnChanged = func(s string) {
		var r, g, b uint8
		l := len(s)
		if l > 7 {
			selector.setColor(selector.tmp)
		} else if _, err := fmt.Sscanf(s, "#%02x%02x%02x", &r, &g, &b); l == 7 && err == nil {
			selector.setColor(color.RGBA{r, g, b, 255})
		} else if _, err := fmt.Sscanf(s, "#%1x%1x%1x", &r, &g, &b); l == 4 && err == nil {
			selector.setColor(color.RGBA{r * 17, g * 17, b * 17, 255})
		}
	}
	return selector
}

func (c *colorSelector) setColor(clr color.Color) {
	c.tmp = c.rect.FillColor
	c.entry.SetText(hexColorString(clr))
	c.rect.FillColor = clr
	c.update(clr)
	reflesh()
}

func hexColorString(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B)
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

func configures1(ts *themeSetting) []fyne.CanvasObject {
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
	return cs
}

func configures2(ts *themeSetting) []fyne.CanvasObject {
	cs := make([]fyne.CanvasObject, 0)
	cs = append(cs, colorConfigure("Scroll bar color", empty(), ts.ScrollBarColor(), ts.SetScrollBarColor)...)
	cs = append(cs, colorConfigure("Shadow color", empty(), ts.ShadowColor(), ts.SetShadowColor)...)
	cs = append(cs, intConfigure("Text size", empty(), ts.TextSize(), ts.SetTextSize)...)
	cs = append(cs, readonlyStringConfigure("Text font", empty(), ts.TextFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold font", empty(), ts.TextBoldFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text italic font", empty(), ts.TextItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text bold italic font", empty(), ts.TextBoldItalicFont().Name())...)
	cs = append(cs, readonlyStringConfigure("Text monospace font", empty(), ts.TextMonospaceFont().Name())...)
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
	text := widget.NewLabel("Text sample")
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

var (
	currentThemeSetting *themeSetting
)

func reflesh() {
	fyne.CurrentApp().Settings().SetTheme(currentThemeSetting)
}

func run(args []string) error {
	a := app.New()
	ts := newThemeSetting()
	currentThemeSetting = ts
	a.Settings().SetTheme(ts)
	w := a.NewWindow("Fyne theme generator")
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewGridLayoutWithColumns(3),
				configures1(ts)...,
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayoutWithColumns(3),
				configures2(ts)...,
			),
		),
	)
	w.ShowAndRun()
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
