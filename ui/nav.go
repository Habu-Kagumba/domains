package ui

import "github.com/jroimartin/gocui"

func (ui *UI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (ui *UI) nextEditor(g *gocui.Gui, v *gocui.View) error {
	return ui.switchEditor(true, false)
}

func (ui *UI) prevEditor(g *gocui.Gui, v *gocui.View) error {
	return ui.switchEditor(false, false)
}

func (ui *UI) wrapEditor(g *gocui.Gui, v *gocui.View) error {
	return ui.switchEditor(true, true)
}

func (ui *UI) switchEditor(forward bool, wrap bool) error {
	var i int

	if forward {
		i = ui.currentView + 1
		if i > len(tabbableViews)-1 {
			if wrap {
				i = 0
			} else {
				return nil
			}
		}
	} else {
		i = ui.currentView - 1
		if i < 0 {
			if wrap {
				i = len(tabbableViews) - 1
			} else {
				return nil
			}
		}
	}

	return ui.setSelectableView(i)
}
