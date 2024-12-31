package components

import (
	"github.com/arrannn/rayout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TextButton(text string, onClick func()) rayout.Widget {
	box := &rayout.Box{
		BackgroundColor: &rayout.GetGlobalTheme().Surface,
		BorderColor:     &rayout.GetGlobalTheme().BorderDark,
		BorderThick:     1,
		AutoWidth:       true,
		AutoHeight:      true,
		PaddingLeft:     5,
		PaddingTop:      5,
		PaddingBottom:   5,
		PaddingRight:    5,
		OnUpdate: func(b *rayout.Box) {
			if b.IsHovered() {
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
					b.BackgroundColor = &rayout.GetGlobalTheme().Background
				} else {
					b.BackgroundColor = &rayout.GetGlobalTheme().Surface
				}
			}
		},
		Child: &rayout.OnClick{
			Action: onClick,
			Child: &rayout.Text{
				Text: text,
				Font: &rayout.GetGlobalTheme().FontBold.Font,
			},
		},
	}
	return box
}
