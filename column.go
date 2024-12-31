package rayout

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Column struct {
	Children []Widget
}

func (c *Column) Update(bounds rl.Rectangle) rl.Rectangle {
	var childrenHeight float32
	var widestChild float32
	for _, w := range c.Children {
		childBounds := w.Update(bounds)
		childrenHeight += childBounds.Height
		if childBounds.Width > widestChild {
			widestChild = childBounds.Width
		}
		bounds.Y += childBounds.Height
	}
	bounds.Height = childrenHeight
	bounds.Width = widestChild
	return bounds
}

func (c *Column) Draw() {
	for _, w := range c.Children {
		w.Draw()
	}
}
