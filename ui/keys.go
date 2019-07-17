package ui

import "github.com/jroimartin/gocui"

type config struct {
	views   *[]string
	key     interface{}
	mod     gocui.Modifier
	handler func(*gocui.Gui, *gocui.View) error
}

func (ui *UI) keyBindings() error {
	var c = []config{
		{
			&views,
			gocui.KeyCtrlQ,
			gocui.ModNone,
			ui.quit,
		},
		{
			&tabbableViews,
			gocui.KeyTab,
			gocui.ModNone,
			ui.wrapEditor,
		},
		{
			&views,
			gocui.KeyCtrlSlash,
			gocui.ModNone,
			ui.inputDomain,
		},
		{
			&[]string{INPUTVIEW},
			gocui.KeyEnter,
			gocui.ModNone,
			ui.search,
		},
		{
			&[]string{INPUTVIEW},
			gocui.KeyCtrlQ,
			gocui.ModNone,
			ui.closeModal,
		},
	}

	for _, binding := range c {
		for _, view := range *binding.views {
			if err := ui.gui.SetKeybinding(view, binding.key, binding.mod, binding.handler); err != nil {
				return err
			}
		}
	}

	return nil
}
