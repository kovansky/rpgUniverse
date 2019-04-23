package visual

import "github.com/therecipe/qt/gui"

type Colors struct {
	ForegroundColor *gui.QColor
	BackgroundColor *gui.QColor
}

func NewColors() Colors {
	return Colors{
		gui.NewQColor3(0, 0, 0, 100),
		gui.NewQColor3(213, 124, 0, 100),
	}
}
