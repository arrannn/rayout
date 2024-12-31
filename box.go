package rayout

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	boundsBox       rl.Rectangle
	borderBox       rl.Rectangle
	ContentBox      rl.Rectangle
	MaxHeight       float32
	MaxWidth        float32
	PaddingTop      float32
	PaddingRight    float32
	PaddingBottom   float32
	PaddingLeft     float32
	MarginTop       float32
	MarginRight     float32
	MarginBottom    float32
	MarginLeft      float32
	AutoWidth       bool
	AutoHeight      bool
	BackgroundColor *color.RGBA
	BorderColor     *color.RGBA
	BorderThick     float32
	Child           Widget
	OnUpdate        func(*Box)
}

func (b *Box) IsHovered() bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), b.borderBox)
}

func (b *Box) Update(bounds rl.Rectangle) rl.Rectangle {
	if b.MaxHeight > 0 && b.MaxHeight < bounds.Height {
		bounds.Height = b.MaxHeight
	}
	if b.MaxWidth > 0 && b.MaxWidth < bounds.Width {
		bounds.Width = b.MaxWidth
	}
	b.boundsBox = bounds
	b.borderBox = ApplyMargin(b.boundsBox, b.MarginTop, b.MarginRight, b.MarginBottom, b.MarginLeft)
	b.ContentBox = ApplyMargin(b.borderBox, b.PaddingTop, b.PaddingRight, b.PaddingBottom, b.PaddingLeft)
	if b.Child == nil {
		return b.boundsBox
	}
	childBox := b.Child.Update(b.ContentBox)
	if b.AutoWidth {
		dif := b.ContentBox.Width - childBox.Width
		b.boundsBox.Width -= dif
		b.borderBox.Width -= dif
		b.ContentBox.Width -= dif
	}
	if b.AutoHeight {
		dif := b.ContentBox.Height - childBox.Height
		b.boundsBox.Height -= dif
		b.borderBox.Height -= dif
		b.ContentBox.Height -= dif
	}

	if b.OnUpdate != nil {
		b.OnUpdate(b)
	}

	return b.boundsBox
}

func (b *Box) Draw() {
	if b.BackgroundColor != nil {
		rl.DrawRectangleRec(b.borderBox, *b.BackgroundColor)
	}
	if b.BorderColor != nil {
		rl.DrawRectangleLinesEx(b.borderBox, b.BorderThick, *b.BorderColor)
	}
	if b.Child != nil {
		b.Child.Draw()
	}
}

func ApplyMargin(rec rl.Rectangle, top, right, bottom, left float32) rl.Rectangle {
	rec.Width -= left + right
	rec.Height -= top + bottom
	rec.X += left
	rec.Y += top
	return rec
}
