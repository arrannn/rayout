package components

import (
	"github.com/arrannn/rayout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputTextComponent struct {
	*rayout.Box
	Value    *string // Pointer to the string value we're modifying
	MaxChars int
	// letterCount  int
	isActive     bool
	frameCounter int
}

func InputText(value *string, maxChars int) *InputTextComponent {
	tb := &InputTextComponent{
		Box: &rayout.Box{
			BackgroundColor: &rayout.GetGlobalTheme().Surface,
			BorderThick:     1,
			PaddingTop:      8,
			PaddingRight:    5,
			PaddingBottom:   8,
			PaddingLeft:     5,
			AutoHeight:      true,
			Child: &rayout.Text{
				Text: *value,
				Font: &rayout.GetGlobalTheme().FontBold.Font,
			},
		},
		Value:    value,
		MaxChars: maxChars,
	}

	// Set up the update handler
	tb.Box.OnUpdate = tb.handleInput

	return tb
}

func (t *InputTextComponent) handleInput(b *rayout.Box) {
	// Check if clicked to activate
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		t.isActive = t.IsHovered()
	}

	if t.isActive {
		b.BorderColor = &rayout.GetGlobalTheme().BorderFocus
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

		b.Child.(*rayout.Text).SetText(*t.Value)

		t.frameCounter++
	} else {
		b.BorderColor = &rayout.GetGlobalTheme().BorderLight
		t.frameCounter = 0
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}

func (t *InputTextComponent) Draw() {
	// First draw the base box
	t.Box.Draw()

	// Draw cursor if active
	if t.isActive && ((t.frameCounter/20)%2) == 0 {
		textWidget := t.Box.Child.(*rayout.Text)
		cursorX := t.Box.ContentBox.X + float32(rl.MeasureTextEx(*textWidget.Font, *t.Value, *textWidget.Size, *textWidget.Spacing).X)
		rl.DrawTextEx(*textWidget.Font, "_", rl.Vector2{X: cursorX, Y: t.Box.ContentBox.Y}, *textWidget.Size, *textWidget.Spacing, *textWidget.Color)
	}
}
