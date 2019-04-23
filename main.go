package main

import (
	"fmt"
	"github.com/kovansky/rpgUniverse/application/visual"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

var (
	app    *widgets.QApplication
	window *widgets.QMainWindow

	palettes = visual.NewPalettes()
)

func main() {
	fmt.Println("Running application")

	app = widgets.NewQApplication(len(os.Args), os.Args)

	window = widgets.NewQMainWindow(nil, core.Qt__Window)
	window.SetWindowTitle("Hello, world!")
	window.SetAutoFillBackground(true)
	window.SetPalette(palettes.WindowPalette)
	window.SetGeometry2(0, 0, 200, 100)

	helloLabel := widgets.NewQLabel2("Hello, world", window, core.Qt__Widget)
	helloLabel.SetPalette(palettes.MainPalette)

	window.Show()
	helloLabel.Show()
	widgets.QApplication_Exec()
}
