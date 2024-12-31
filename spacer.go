package rayout

import rl "github.com/gen2brain/raylib-go/raylib"

type Spacer struct {
	Width  float32
	Height float32
}

func (s *Spacer) Update(bounds rl.Rectangle) rl.Rectangle {
	bounds.Width = s.Width
	bounds.Height = s.Height
	return bounds
}

func (s *Spacer) Draw() {}
