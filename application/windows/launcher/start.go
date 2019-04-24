package launcher

import (
	"github.com/kovansky/rpgUniverse/application"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	window *widgets.QMainWindow
)

func RunStart() {
	window = widgets.NewQMainWindow(nil, core.Qt__Window)
	window.SetWindowTitle(application.Language.App.Title)
	window.SetAutoFillBackground(true)
	window.SetPalette(application.Palettes.WindowPalette)
	window.SetGeometry2(0, 0, 200, 100)

	helloLabel := widgets.NewQLabel2(application.Language.App.Title, window, core.Qt__Widget)
	helloLabel.SetPalette(application.Palettes.MainPalette)

	window.Show()
	helloLabel.Show()
}
