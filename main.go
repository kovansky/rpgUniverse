package main

import (
	"fmt"
	"github.com/kovansky/rpgUniverse/application/windows/launcher"
	"github.com/therecipe/qt/widgets"
	"os"
)

var (
	app *widgets.QApplication
)

func main() {
	fmt.Println("Running application")

	app = widgets.NewQApplication(len(os.Args), os.Args)

	window = widgets.NewQMainWindow(nil, core.Qt__Window)
	window.SetWindowTitle(Language.App.Title)
	window.SetAutoFillBackground(true)
	window.SetPalette(palettes.WindowPalette)
	window.SetGeometry2(0, 0, 200, 100)

	helloLabel := widgets.NewQLabel2(Language.App.Title, window, core.Qt__Widget)
	helloLabel.SetPalette(palettes.MainPalette)

	window.Show()
	helloLabel.Show()
	widgets.QApplication_Exec()
}
