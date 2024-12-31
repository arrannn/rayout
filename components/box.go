package components

import "github.com/arrannn/rayout"

func PaddedBox(padding float32, child rayout.Widget) *rayout.Box {
	return &rayout.Box{
		Child:         child,
		PaddingTop:    padding,
		PaddingBottom: padding,
		PaddingRight:  padding,
		PaddingLeft:   padding,
	}
}
