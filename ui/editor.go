package ui

import "github.com/jroimartin/gocui"

type editor struct {
	gocuiEditor gocui.Editor
}

var ed editor

func (e *editor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch key {
	case gocui.KeyEnter:
		return
	case gocui.KeyArrowRight:
		x, _ := v.Cursor()
		if x >= len(v.ViewBuffer())-1 {
			return
		}
	case gocui.KeyHome:
		v.SetCursor(0, 0)
		return
	case gocui.KeyEnd:
		v.SetCursor(len(v.ViewBuffer())-2, 0)
		return
	}

	e.gocuiEditor.Edit(v, key, ch, mod)
}
