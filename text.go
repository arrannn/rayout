package rayout

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Text    string
	pos     rl.Vector2
	Font    *rl.Font
	Size    *float32
	Spacing *float32
	Color   *rl.Color
}

func NewText(text string) *Text {
	return &Text{Text: text}
}

func (t *Text) SetText(text string) {
	t.Text = text
}

func (t *Text) Update(bounds rl.Rectangle) rl.Rectangle {
	if t.Font == nil {
		t.Font = &globalTheme.FontRegular.Font
	}
	if t.Size == nil {
		t.Size = &globalTheme.FontRegular.Size
	}
	if t.Spacing == nil {
		t.Spacing = &globalTheme.FontRegular.Spacing
	}
	if t.Color == nil {
		t.Color = &globalTheme.TextPrimary
	}

	textDimensions := rl.MeasureTextEx(*t.Font, t.Text, *t.Size, *t.Spacing)
	bounds.Width = textDimensions.X
	bounds.Height = textDimensions.Y
	t.pos = rl.Vector2{X: bounds.X, Y: bounds.Y}
	return bounds
}

func (t *Text) Draw() {
	rl.DrawTextEx(
		*t.Font,
		t.Text,
		t.pos,
		*t.Size,
		*t.Spacing,
		*t.Color,
	)
}
