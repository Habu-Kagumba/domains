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
		// All Views
		{
			&views,
			gocui.KeyCtrlQ,
			gocui.ModNone,
			ui.quit,
		},
		{
			&views,
			gocui.KeyCtrlSlash,
			gocui.ModNone,
			ui.inputDomain,
		},
		{
			&views,
			gocui.KeyCtrlH,
			gocui.ModNone,
			ui.help,
		},
		// Tabbable Views
		{
			&tabbableViews,
			gocui.KeyTab,
			gocui.ModNone,
			ui.wrapEditor,
		},
		// INPUTVIEW
		{
			&[]string{INPUTVIEW},
			gocui.KeyEnter,
			gocui.ModNone,
			ui.search,
		},
		{
			&[]string{INPUTVIEW},
			gocui.KeyEsc,
			gocui.ModNone,
			ui.closeModal,
		},
		// HELPVIEW
		{
			&[]string{HELPVIEW},
			gocui.KeyEsc,
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
