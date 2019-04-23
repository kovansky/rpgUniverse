package visual

import "github.com/therecipe/qt/gui"

type Palettes struct {
	WindowPalette *gui.QPalette
	MainPalette   *gui.QPalette
}

func NewPalettes() Palettes {
	colors := NewColors()

	palettes := Palettes{
		gui.NewQPalette(),
		gui.NewQPalette(),
	}

	palettes.WindowPalette.SetColor2(gui.QPalette__Background, colors.BackgroundColor)
	palettes.MainPalette.SetColor2(gui.QPalette__Foreground, colors.ForegroundColor)

	return palettes
}
