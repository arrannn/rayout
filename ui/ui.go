package ui

import (
	"github.com/arrannn/rayout/themes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var globalTheme *themes.Theme

func SetGlobalTheme(theme *themes.Theme) {
	globalTheme = theme
}

func GetGlobalTheme() *themes.Theme {
	return globalTheme
}

type Widget interface {
	Update(bounds rl.Rectangle) rl.Rectangle
	Draw()
}
