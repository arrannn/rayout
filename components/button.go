package components

import (
	"github.com/arrannn/rayout/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TextButton(text string, onClick func()) ui.Widget {
	box := &ui.Box{
		BackgroundColor: &ui.GetGlobalTheme().Surface,
		BorderColor:     &ui.GetGlobalTheme().BorderDark,
		BorderThick:     1,
		AutoWidth:       true,
		AutoHeight:      true,
		PaddingLeft:     5,
		PaddingTop:      5,
		PaddingBottom:   5,
		PaddingRight:    5,
		OnUpdate: func(b *ui.Box) {
			if b.IsHovered() {
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
					b.BackgroundColor = &ui.GetGlobalTheme().Background
				} else {
					b.BackgroundColor = &ui.GetGlobalTheme().Surface
				}
			}
		},
		Child: &ui.OnClick{
			Action: onClick,
			Child: &ui.Text{
				Text: text,
				Font: &ui.GetGlobalTheme().FontBold.Font,
			},
		},
	}
	return box
}
