package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

const (
	INPUTVIEW            = "input"
	PREVIOUSSEARCHESVIEW = "previous"
	AVAILABLEVIEW        = "available"
	SUGGESTIONSVIEW      = "suggestions"
	FOOTERVIEW           = "footer"
)

type properties struct {
	title    string
	content  string
	x1       float64
	y1       float64
	x2       float64
	y2       float64
	editable bool
	modal    bool
}

var viewsProps = map[string]properties{
	INPUTVIEW: {
		title:    "Enter a domain name",
		content:  "",
		editable: true,
		modal:    true,
	},
	PREVIOUSSEARCHESVIEW: {
		title:    "Previous Searches",
		content:  "",
		x1:       0.0,
		y1:       0.0,
		x2:       0.3,
		y2:       0.9,
		editable: false,
		modal:    false,
	},
	AVAILABLEVIEW: {
		title:    "Available Domains",
		content:  "",
		x1:       0.31,
		y1:       0,
		x2:       1,
		y2:       0.5,
		editable: false,
		modal:    false,
	},
	SUGGESTIONSVIEW: {
		title:    "Suggestions",
		content:  "",
		x1:       0.31,
		y1:       0.51,
		x2:       1,
		y2:       0.9,
		editable: false,
		modal:    false,
	},
	FOOTERVIEW: {
		title:    "Console",
		content:  "",
		x1:       0,
		y1:       0.91,
		x2:       1,
		y2:       1,
		editable: false,
		modal:    false,
	},
}

var views = []string{
	PREVIOUSSEARCHESVIEW,
	AVAILABLEVIEW,
	SUGGESTIONSVIEW,
	FOOTERVIEW,
}

var tabbableViews = []string{
	PREVIOUSSEARCHESVIEW,
	AVAILABLEVIEW,
	SUGGESTIONSVIEW,
}

func (ui *UI) Layout(g *gocui.Gui) error {
	for _, v := range views {
		if err := ui.initView(v); err != nil {
			return err
		}
	}

	if ui.currentView == -1 {
		ui.currentView = 0
		ui.setSelectableView(ui.currentView)
	}

	return nil
}

func (ui *UI) initView(name string) error {
	maxX, maxY := ui.gui.Size()

	vp := viewsProps[name]

	if vp.modal {
		return nil
	}

	x1 := int(vp.x1 * float64(maxX))
	y1 := int(vp.y1 * float64(maxY))
	x2 := int(vp.x2*float64(maxX)) - 1
	y2 := int(vp.y2*float64(maxY)) - 1

	return ui.createView(name, x1, y1, x2, y2)
}

func (ui *UI) createView(name string, x1, y1, x2, y2 int) error {
	if v, err := ui.gui.SetView(name, x1, y1, x2, y2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		vp := viewsProps[name]
		v.Title = vp.title
		v.Editable = vp.editable
		if vp.editable {
			v.Editor = &ed
		}

		ui.writeView(name, vp.content)
	}

	return nil
}

func (ui *UI) setSelectableView(id int) error {
	if err := ui.setView(views[id]); err != nil {
		return err
	}
	ui.currentView = id

	return nil
}

func (ui *UI) setView(name string) error {
	if _, err := ui.gui.SetCurrentView(name); err != nil {
		return err
	}

	return nil
}

func (ui *UI) writeView(name string, content string) {
	v, _ := ui.gui.View(name)
	v.Clear()
	fmt.Fprint(v, content)
	v.SetCursor(len(content), 0)
}

func (ui *UI) clearView(name string) {
	v, _ := ui.gui.View(name)
	v.Clear()
}

func (ui *UI) closeView(name string) {
	ui.gui.DeleteView(name)
	ui.setView(PREVIOUSSEARCHESVIEW)
}

func (ui *UI) showModal(name string, width, height float64) {
	vp := viewsProps[name]
	viewsProps[name] = vp

	maxX, maxY := ui.gui.Size()

	modalWidth := int(float64(maxX) * width)
	modalHeight := int(float64(maxY) * height)

	x1 := (maxX - modalWidth) / 2
	x2 := x1 + modalWidth
	y1 := (maxY - modalHeight) / 3
	y2 := y1 + modalHeight

	ui.createView(name, x1, y1, x2, y2)
	ui.setView(name)
}

func (ui *UI) writeConsole(content string, isError bool) {
	if isError {
		content = decorate(content, "red")
	}

	ui.writeView(FOOTERVIEW, content)
}

func (ui *UI) parseDomain(name string) string {
	v, _ := ui.gui.View(name)
	parts := strings.Fields(v.Buffer())

	m := make(map[string]bool)
	unique := make([]string, 0, len(parts))
	for _, part := range parts {
		if ok := m[part]; !ok {
			unique = append(unique, part)
			m[part] = true
		}
	}

	return strings.Join(unique[:], "")
}
