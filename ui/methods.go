package ui

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Habu-Kagumba/domains/suggestions"
	"github.com/jroimartin/gocui"
	"github.com/ttacon/chalk"
)

// TODO refactor this
func handleErrors(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (ui *UI) closeModal(g *gocui.Gui, v *gocui.View) error {
	ui.closeView(v.Name())
	return nil
}

func (ui *UI) inputDomain(g *gocui.Gui, v *gocui.View) error {
	ui.showModal(INPUTVIEW, 0.7, 0.05)
	return nil
}

func getAbsPath(p string) (absPath string) {
	absPath, err := filepath.Abs(p)
	handleErrors(err)
	return
}

func getFile(p string) (f *os.File) {
	f, err := os.Open(p)
	handleErrors(err)
	return
}

func (ui *UI) search(g *gocui.Gui, v *gocui.View) error {
	ui.writeConsole("Searching...", false, false)
	ui.clearView(SUGGESTIONSVIEW)

	p := getAbsPath("suggestions/extensions/pref-ix-suff.json")
	f := getFile(p)

	defer func() {
		err := f.Close()
		handleErrors(err)
	}()

	e := suggestions.LoadNameExtensions(f)
	s := suggestions.Suggestions(ui.parseDomain(INPUTVIEW), e)

	ui.closeView(INPUTVIEW)
	ui.writeConsole("Done", false, true)
	ui.writeView(SUGGESTIONSVIEW, strings.Join(s[:], "\n"))

	return nil
}

func (ui *UI) help(g *gocui.Gui, v *gocui.View) error {
	ui.showModal(HELPVIEW, 0.7, 0.1)
	return nil
}

func decorate(s string, color string) string {
	switch color {
	case "green":
		s = chalk.Green.Color(s)
	case "red":
		s = chalk.Red.Color(s)
	case "cyan":
		s = chalk.Cyan.Color(s)
	default:
		return s
	}

	return s
}
