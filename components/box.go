package components

import "github.com/arrannn/rayout/ui"

func PaddedBox(padding float32, child ui.Widget) *ui.Box {
	return &ui.Box{
		Child:         child,
		PaddingTop:    padding,
		PaddingBottom: padding,
		PaddingRight:  padding,
		PaddingLeft:   padding,
	}
}
