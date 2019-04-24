package application

import (
	"encoding/json"
	"github.com/kovansky/rpgUniverse/application/schemas/languages"
	"github.com/kovansky/rpgUniverse/application/schemas/settings"
	"github.com/kovansky/rpgUniverse/application/visual"
	"io/ioutil"
	"os"
)

var (
	Config   *settings.Config
	Language *languages.Language

	Palettes = visual.NewPalettes()
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
