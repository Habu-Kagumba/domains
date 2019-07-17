package main

import (
	"github.com/Habu-Kagumba/domains/ui"
)

func main() {
	u := ui.NewUI()
	defer u.Close()

	u.Loop()
}
