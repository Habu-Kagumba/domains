package ui

import (
	"log"

	"github.com/jroimartin/gocui"
)

// UI struct
type UI struct {
	gui         *gocui.Gui
	currentView int
}

// NewUI creates a new gocui UI
func NewUI() *UI {
	var err error
	ui := new(UI)
	ui.gui, err = gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}

	ui.initGui()

	return ui
}

// Loop runs MainLoop until error is returned
func (ui *UI) Loop() {
	if err := ui.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// Close exits GUI
func (ui *UI) Close() {
	ui.gui.Close()
}

func (ui *UI) initGui() {
	ed = editor{gocui.DefaultEditor}

	// Defaults
	ui.gui.Cursor = true
	ui.gui.BgColor = gocui.ColorDefault
	ui.gui.FgColor = gocui.Attribute(15 + 1)

	// Set Layout function
	ui.gui.SetManagerFunc(ui.Layout)

	ui.currentView = -1

	ui.keyBindings()
}
