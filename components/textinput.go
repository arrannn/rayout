package components

import (
	"github.com/arrannn/rayout/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextInput struct {
	*ui.Box
	Value    *string // Pointer to the string value we're modifying
	MaxChars int
	// letterCount  int
	isActive     bool
	frameCounter int
}

func NewTextInput(value *string, maxChars int) *TextInput {
	tb := &TextInput{
		Box: &ui.Box{
			BackgroundColor: &ui.GetGlobalTheme().Surface,
			BorderThick:     1,
			PaddingTop:      8,
			PaddingRight:    5,
			PaddingBottom:   8,
			PaddingLeft:     5,
			AutoHeight:      true,
			Child: &ui.Text{
				Text: *value,
				Font: &ui.GetGlobalTheme().FontBold.Font,
			},
		},
		Value:    value,
		MaxChars: maxChars,
	}

	// Set up the update handler
	tb.Box.OnUpdate = tb.handleInput

	return tb
}

func (t *TextInput) handleInput(b *ui.Box) {
	// Check if clicked to activate
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		t.isActive = t.IsHovered()
	}

	if t.isActive {
		b.BorderColor = &ui.GetGlobalTheme().BorderFocus
		rl.SetMouseCursor(rl.MouseCursorIBeam)

		// Handle character input
		key := rl.GetCharPressed()
		for key > 0 {
			if key >= 32 && key <= 125 && len(*t.Value) < t.MaxChars {
				*t.Value += string(key)
			}
			key = rl.GetCharPressed()
		}

		// Handle backspace
		if rl.IsKeyPressed(rl.KeyBackspace) && len(*t.Value) > 0 {
			*t.Value = (*t.Value)[:len(*t.Value)-1]
		}

		b.Child.(*ui.Text).SetText(*t.Value)

		t.frameCounter++
	} else {
		b.BorderColor = &ui.GetGlobalTheme().BorderLight
		t.frameCounter = 0
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}

func (t *TextInput) Draw() {
	// First draw the base box
	t.Box.Draw()

	// Draw cursor if active
	if t.isActive && ((t.frameCounter/20)%2) == 0 {
		textWidget := t.Box.Child.(*ui.Text)
		cursorX := t.Box.ContentBox.X + float32(rl.MeasureTextEx(*textWidget.Font, *t.Value, *textWidget.Size, *textWidget.Spacing).X)
		rl.DrawTextEx(*textWidget.Font, "_", rl.Vector2{X: cursorX, Y: t.Box.ContentBox.Y}, *textWidget.Size, *textWidget.Spacing, *textWidget.Color)
	}
}
