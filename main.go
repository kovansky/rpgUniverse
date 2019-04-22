package main

import (
	"fmt"
	"github.com/kovansky/rpgUniverse/application"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

var (
	app    *widgets.QApplication
	window *widgets.QMainWindow

	winPalette  *gui.QPalette
	mainPalette *gui.QPalette
)

func main() {
	fmt.Println("Running application")

	winPalette = gui.NewQPalette()
	winPalette.SetColor2(gui.QPalette__Background, application.ColorBackground)

	mainPalette = gui.NewQPalette()
	mainPalette.SetColor2(gui.QPalette__Foreground, application.ColorForeground)

	app = widgets.NewQApplication(len(os.Args), os.Args)

	window = widgets.NewQMainWindow(nil, core.Qt__Window)
	window.SetWindowTitle("Hello, world!")
	window.SetAutoFillBackground(true)
	window.SetPalette(winPalette)
	window.SetGeometry2(0, 0, 200, 100)

	helloLabel := widgets.NewQLabel2("Hello, world", window, core.Qt__Widget)

	window.Show()
	helloLabel.Show()
	widgets.QApplication_Exec()
}
