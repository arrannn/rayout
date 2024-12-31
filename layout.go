package rayout

import (
	"github.com/arrannn/rayout/themes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout struct {
	rect    rl.Rectangle
	Content Widget
}

func NewLayout(rect rl.Rectangle) *Layout {
	if globalTheme == nil {
		SetGlobalTheme(themes.GorayTheme())
	}
	return &Layout{rect: rect}
}

func (l *Layout) Update() {
	l.Content.Update(l.rect)
}

func (l *Layout) Draw() {
	l.Content.Draw()
}

func (l *Layout) UpdateAndDraw() {
	l.Update()
	l.Draw()
}
