package main

import (
	"encoding/json"
	"fmt"
	"github.com/kovansky/rpgUniverse/application/schemas/languages"
	"github.com/kovansky/rpgUniverse/application/schemas/settings"
	"github.com/kovansky/rpgUniverse/application/visual"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
)

var (
	Config   *settings.Config
	Language *languages.Language

	app    *widgets.QApplication
	window *widgets.QMainWindow

	palettes = visual.NewPalettes()
)

/*
Open config file
*/
func init() {
	configFile, err := os.Open("settings/config.json")

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &Config)

	if err != nil {
		panic(err)
	}

	defer configFile.Close()
}

/*
Open language file
*/
func init() {
	languageFile, err := os.Open("settings/languages/" + Config.General.Language + ".json")

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(languageFile)

	err = json.Unmarshal(byteValue, &Language)

	if err != nil {
		panic(err)
	}

	defer languageFile.Close()
}

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
