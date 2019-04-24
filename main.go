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

	launcher.RunStart()

	widgets.QApplication_Exec()
}
